package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
	"restapi-with-opentelemetry/config"
	"restapi-with-opentelemetry/pkg/observability"
)

type server struct {
	router  http.Handler
	address string

	stopPusherMetricsFn observability.StopPusherFunc
	shutdownExporterFn  observability.ShutDownFunc
}

func NewServer() *server {
	//setup metricsPusher
	stopPusher := observability.InitMetricProvider()

	//setup tracer exporter
	shutdownExporter := observability.InitTracerProvider()

	//from wire gen //dependency injector
	router := InitializedServerRouter()

	return &server{
		router:              router,
		address:             config.ServiceAddress,
		stopPusherMetricsFn: stopPusher,
		shutdownExporterFn:  shutdownExporter,
	}
}

func (s *server) Run() {
	srv := &http.Server{
		Addr:    s.address,
		Handler: s.router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("unable to create http listener")
		}
	}()

	log.Debug().Msgf("web server started on %v", s.address)

	<-done
	log.Debug().Msg("web server has been stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctxShutDown); err != nil {
		log.Err(err).Msg("failed to shutdown web server")
	}
	log.Debug().Msg("web server exited properly")

	if err := s.shutdownExporterFn(ctxShutDown); err != nil {
		log.Err(err).Msg("failed to shutdown tracer exporter")
	}
	log.Debug().Msg("tracer exporter exited properly")

	if err := s.stopPusherMetricsFn(ctxShutDown); err != nil {
		log.Err(err).Msg("failed to stoping pusher metrics")
	}
	log.Debug().Msg("metrics pusher exited properly")
	log.Debug().Msg("done, server exited properly :)")
}
