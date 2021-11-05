package config

import (
	"context"
	"fmt"
	"reflect"

	gofig "github.com/gookit/config/v2"
)

// Log defines the available logging configuration.
type Log struct {
	Level  string `mapstructure:"level"`
	Pretty bool   `mapstructure:"pretty"`
	Color  bool   `mapstructure:"color"`
	File   string `mapstructure:"file"`
}

// Debug defines the available debug configuration.
type Debug struct {
	Addr   string `mapstructure:"addr"`
	Token  string `mapstructure:"token"`
	Pprof  bool   `mapstructure:"pprof"`
	Zpages bool   `mapstructure:"zpages"`
}

// HTTP defines the available http configuration.
type HTTP struct {
	Addr      string `mapstructure:"addr"`
	Root      string `mapstructure:"root"`
	Namespace string `mapstructure:"namespace"`
	CacheTTL  int    `mapstructure:"cache_ttl"`
}

// Tracing defines the available tracing configuration.
type Tracing struct {
	Enabled   bool   `mapstructure:"enabled"`
	Type      string `mapstructure:"type"`
	Endpoint  string `mapstructure:"endpoint"`
	Collector string `mapstructure:"collector"`
	Service   string `mapstructure:"service"`
}

// Asset defines the available asset configuration.
type Asset struct {
	Path string `mapstructure:"path"`
}

// WebConfig defines the available web configuration for a dynamically rendered config.json.
type WebConfig struct {
	Server        string                 `json:"server,omitempty" mapstructure:"server"`
	Theme         string                 `json:"theme,omitempty" mapstructure:"theme"`
	Version       string                 `json:"version,omitempty" mapstructure:"version"` // TODO what is version used for?
	OpenIDConnect OIDC                   `json:"openIdConnect,omitempty" mapstructure:"oids"`
	Apps          []string               `json:"apps" mapstructure:"apps"` // TODO add nil as empty when https://go-review.googlesource.com/c/go/+/205897/ is released
	ExternalApps  []ExternalApp          `json:"external_apps,omitempty" mapstructure:"external_apps"`
	Options       map[string]interface{} `json:"options,omitempty" mapstructure:"options"`
}

// OIDC defines the available oidc configuration
type OIDC struct {
	MetadataURL  string `json:"metadata_url,omitempty" mapstructure:"metadata_url"`
	Authority    string `json:"authority,omitempty" mapstructure:"authority"`
	ClientID     string `json:"client_id,omitempty" mapstructure:"client_id"`
	ResponseType string `json:"response_type,omitempty" mapstructure:"response_type"`
	Scope        string `json:"scope,omitempty" mapstructure:"scope"`
}

// ExternalApp defines an external web app.
// {
//	"name": "hello",
//	"path": "http://localhost:9105/hello.js",
//	  "config": {
//	    "url": "http://localhost:9105"
//	  }
//  }
type ExternalApp struct {
	ID   string `json:"id,omitempty" mapstructure:"id"`
	Path string `json:"path,omitempty" mapstructure:"path"`
	// Config is completely dynamic, because it depends on the extension
	Config map[string]interface{} `json:"config,omitempty" mapstructure:"config"`
}

// ExternalAppConfig defines an external web app configuration.
type ExternalAppConfig struct {
	URL string `json:"url,omitempty" mapstructure:"url"`
}

// Web defines the available web configuration.
type Web struct {
	Path        string    `mapstructure:"path"`
	ThemeServer string    `mapstructure:"theme_server"` // used to build Theme in WebConfig
	ThemePath   string    `mapstructure:"theme_path"`   // used to build Theme in WebConfig
	Config      WebConfig `mapstructure:"config"`
}

// Config combines all available configuration parts.
type Config struct {
	File    string  `mapstructure:"file"`
	Log     Log     `mapstructure:"log"`
	Debug   Debug   `mapstructure:"debug"`
	HTTP    HTTP    `mapstructure:"http"`
	Tracing Tracing `mapstructure:"tracing"`
	Asset   Asset   `mapstructure:"asset"`
	Web     Web     `mapstructure:"web"`

	Context    context.Context
	Supervised bool
}

// New initializes a new configuration with or without defaults.
func New() *Config {
	return &Config{}
}

func DefaultConfig() *Config {
	return &Config{
		Log: Log{},
		Debug: Debug{
			Addr:   "127.0.0.1:9104",
			Token:  "",
			Pprof:  false,
			Zpages: false,
		},
		HTTP: HTTP{
			Addr:      "127.0.0.1:9100",
			Root:      "/",
			Namespace: "com.owncloud.web",
			CacheTTL:  604800, // 7 days
		},
		Tracing: Tracing{
			Enabled:   false,
			Type:      "jaeger",
			Endpoint:  "",
			Collector: "",
			Service:   "web",
		},
		Asset: Asset{
			Path: "",
		},
		Web: Web{
			Path:        "",
			ThemeServer: "https://localhost:9200",
			ThemePath:   "/themes/owncloud/theme.json",
			Config: WebConfig{
				Server:  "https://localhost:9200",
				Theme:   "",
				Version: "0.1.0",
				OpenIDConnect: OIDC{
					MetadataURL:  "",
					Authority:    "https://localhost:9200",
					ClientID:     "web",
					ResponseType: "code",
					Scope:        "openid profile email",
				},
				Apps: []string{"files", "search", "media-viewer", "external"},
			},
		},
	}
}

// GetEnv fetches a list of known env variables for this extension. It is to be used by gookit, as it provides a list
// with all the environment variables an extension supports.
func GetEnv() []string {
	var r = make([]string, len(structMappings(&Config{})))
	for i := range structMappings(&Config{}) {
		r = append(r, structMappings(&Config{})[i].EnvVars...)
	}

	return r
}

// UnmapEnv loads values from the gooconf.Config argument and sets them in the expected destination.
func (c *Config) UnmapEnv(gooconf *gofig.Config) error {
	vals := structMappings(c)
	for i := range vals {
		for j := range vals[i].EnvVars {
			// we need to guard against v != "" because this is the condition that checks that the value is set from the environment.
			// the `ok` guard is not enough, apparently.
			if v, ok := gooconf.GetValue(vals[i].EnvVars[j]); ok && v != "" {

				// get the destination type from destination
				switch reflect.ValueOf(vals[i].Destination).Type().String() {
				case "*bool":
					r := gooconf.Bool(vals[i].EnvVars[j])
					*vals[i].Destination.(*bool) = r
				case "*string":
					r := gooconf.String(vals[i].EnvVars[j])
					*vals[i].Destination.(*string) = r
				case "*int":
					r := gooconf.Int(vals[i].EnvVars[j])
					*vals[i].Destination.(*int) = r
				case "*float64":
					// defaults to float64
					r := gooconf.Float(vals[i].EnvVars[j])
					*vals[i].Destination.(*float64) = r
				default:
					// it is unlikely we will ever get here. Let this serve more as a runtime check for when debugging.
					return fmt.Errorf("invalid type for env var: `%v`", vals[i].EnvVars[j])
				}
			}
		}
	}

	return nil
}
