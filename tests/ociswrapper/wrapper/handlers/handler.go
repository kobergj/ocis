package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"ociswrapper/common"
	"ociswrapper/ocis"
	"os"
)

type BasicResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
type CommandResponse struct {
	*BasicResponse
	ExitCode int `json:"exitCode"`
}

func parseJsonBody(reqBody io.ReadCloser) (map[string]any, error) {
	body, _ := io.ReadAll(reqBody)

	if len(body) == 0 || !json.Valid(body) {
		return nil, errors.New("invalid json data")
	}

	var bodyMap map[string]any
	json.Unmarshal(body, &bodyMap)

	return bodyMap, nil
}

func sendResponse(res http.ResponseWriter, statusCode int, message string) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)

	var status string
	if statusCode == http.StatusOK {
		status = "OK"
	} else {
		status = "ERROR"
	}

	resBody := BasicResponse{
		Status:  status,
		Message: message,
	}

	jsonResponse, _ := json.Marshal(resBody)
	res.Write(jsonResponse)
}

func sendCmdResponse(res http.ResponseWriter, exitCode int, message string) {
	resBody := CommandResponse{
		BasicResponse: &BasicResponse{
			Message: message,
		},
		ExitCode: exitCode,
	}

	if exitCode == 0 {
		resBody.BasicResponse.Status = "OK"
	} else {
		resBody.BasicResponse.Status = "ERROR"
	}

	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")

	jsonResponse, _ := json.Marshal(resBody)
	res.Write(jsonResponse)
}

func SetEnvHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPut {
		sendResponse(res, http.StatusMethodNotAllowed, "")
		return
	}

	envBody, err := parseJsonBody(req.Body)
	if err != nil {
		sendResponse(res, http.StatusMethodNotAllowed, "Invalid json body")
		return
	}

	var envMap []string
	for key, value := range envBody {
		envMap = append(envMap, fmt.Sprintf("%s=%v", key, value))
	}
	ocis.EnvConfigs = append(ocis.EnvConfigs, envMap...)

	var message string

	success, _ := ocis.Restart(ocis.EnvConfigs)
	if success {
		message = "oCIS configured successfully"
		sendResponse(res, http.StatusOK, message)
		return
	}

	message = "Failed to restart oCIS with new configuration"
	sendResponse(res, http.StatusInternalServerError, message)
}

func RollbackHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		sendResponse(res, http.StatusMethodNotAllowed, "")
		return
	}

	var message string
	ocis.EnvConfigs = []string{}
	success, _ := ocis.Restart(os.Environ())
	if success {
		message = "oCIS configuration rolled back successfully"
		sendResponse(res, http.StatusOK, message)
		return
	}

	message = "Failed to restart oCIS with initial configuration"
	sendResponse(res, http.StatusInternalServerError, message)
}

func StopOcisHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		sendResponse(res, http.StatusMethodNotAllowed, "")
		return
	}

	success, message := ocis.Stop()
	if success {
		sendResponse(res, http.StatusOK, message)
		return
	}

	sendResponse(res, http.StatusInternalServerError, message)
}

func StartOcisHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		sendResponse(res, http.StatusMethodNotAllowed, "")
		return
	}

	if ocis.IsOcisRunning() {
		sendResponse(res, http.StatusConflict, "oCIS server is already running")
		return
	}

	common.Wg.Add(1)
	go ocis.Start(nil)

	success, message := ocis.WaitForConnection()
	if success {
		sendResponse(res, http.StatusOK, message)
		return
	}

	sendResponse(res, http.StatusInternalServerError, message)
}

func CommandHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		sendResponse(res, http.StatusMethodNotAllowed, "")
		return
	}

	if req.Body == nil {
		sendResponse(res, http.StatusBadRequest, "Body is missing")
		return
	}

	body, err := parseJsonBody(req.Body)
	if err != nil {
		sendResponse(res, http.StatusBadRequest, "Invalid json body")
		return
	}
	if _, ok := body["command"]; !ok {
		sendResponse(res, http.StatusBadRequest, "Command is missing")
		return
	}

	command := body["command"].(string)

	stdIn := []string{}
	if _, ok := body["inputs"]; ok {
		if inputs, ok := body["inputs"].([]interface{}); ok {
			for _, input := range inputs {
				if _, ok := input.(string); ok {
					stdIn = append(stdIn, input.(string))
				} else {
					sendResponse(res, http.StatusBadRequest, "Invalid input data. Expected string")
					return
				}
			}
		}
	}

	exitCode, output := ocis.RunCommand(command, stdIn)
	sendCmdResponse(res, exitCode, output)
}
