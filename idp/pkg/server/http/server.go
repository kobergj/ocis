package http

import (
	"crypto/tls"
	"os"

	"github.com/owncloud/ocis/idp/pkg/crypto"
	svc "github.com/owncloud/ocis/idp/pkg/service/v0"
	"github.com/owncloud/ocis/ocis-pkg/middleware"
	"github.com/owncloud/ocis/ocis-pkg/service/http"
)

// Server initializes the http service and server.
func Server(opts ...Option) (http.Service, error) {
	options := newOptions(opts...)

	var tlsConfig *tls.Config
	if options.Config.HTTP.TLS {
		if options.Config.HTTP.TLSCert == "" || options.Config.HTTP.TLSKey == "" {
			_, certErr := os.Stat("./data/proxy/server.crt")
			_, keyErr := os.Stat("./data/proxy/server.key")

			if os.IsNotExist(certErr) || os.IsNotExist(keyErr) {
				options.Logger.Info().Msgf("Generating certs")
				if err := crypto.GenCert(options.Logger); err != nil {
					options.Logger.Fatal().Err(err).Msg("Could not setup TLS")
					os.Exit(1)
				}
			}

			options.Config.HTTP.TLSCert = "./data/proxy/server.crt"
			options.Config.HTTP.TLSKey = "./data/proxy/server.key"
		}

		cer, err := tls.LoadX509KeyPair(options.Config.HTTP.TLSCert, options.Config.HTTP.TLSKey)
		if err != nil {
			options.Logger.Fatal().Err(err).Msg("Could not setup TLS")
			os.Exit(1)
		}

		tlsConfig = &tls.Config{Certificates: []tls.Certificate{cer}}
	}

	service := http.NewService(
		http.Logger(options.Logger),
		http.Namespace(options.Config.Service.Namespace),
		http.Name(options.Config.Service.Name),
		http.Version(options.Config.Service.Version),
		http.Address(options.Config.HTTP.Addr),
		http.Context(options.Context),
		http.Flags(options.Flags...),
		http.TLSConfig(tlsConfig),
	)

	handle := svc.NewService(
		svc.Logger(options.Logger),
		svc.Config(options.Config),
		svc.Middleware(
			middleware.RealIP,
			middleware.RequestID,
			middleware.NoCache,
			middleware.Cors,
			middleware.Secure,
			middleware.Version(
				options.Config.Service.Name,
				options.Config.Service.Version,
			),
			middleware.Logger(
				options.Logger,
			),
		),
	)

	{
		handle = svc.NewTracing(handle)
		handle = svc.NewInstrument(handle, options.Metrics)
		handle = svc.NewLogging(handle, options.Logger)
	}

	service.Handle(
		"/",
		handle,
	)

	service.Init()
	return service, nil
}
