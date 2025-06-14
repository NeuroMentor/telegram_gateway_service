package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"telegram_gateway_service/internal/config"
	"telegram_gateway_service/internal/runner"
	http_server "telegram_gateway_service/internal/server/http"
	"telegram_gateway_service/internal/shared/middleware"
	"telegram_gateway_service/internal/shared/response"
	transperr "telegram_gateway_service/internal/shared/transport_error"
	"telegram_gateway_service/internal/webhook"
	"telegram_gateway_service/pkg/logger"

	"github.com/go-openapi/strfmt"
	"github.com/gorilla/mux"
)

func main() {
	var ctx = context.Background()

	log, err := logger.NewLogger()
	if err != nil {
		panic("failed to create logger: " + err.Error())
	}
	cfg := config.MustConfig(log)

	var (
		mid              = middleware.NewMiddleware(log)
		httpResp         = response.NewHTTPResponse(log, true)
		convert          = transperr.NewErrorConverter()
		validationFormat = strfmt.NewFormats()

		router = mux.NewRouter()
	)

	initBusinessLogic(
		router,

		mid,
		httpResp,
		convert,
		validationFormat,

		log,
	)
	httpServer := http_server.NewServer(&cfg.Server, router)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("HTTP server failed: %v", err)
		}
	}()

	log.Infof("server listening on port [%d] | Env %s", cfg.Server.Port, cfg.Env)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	err = httpServer.Shutdown(ctx)
	if err != nil {
		log.Panicf("error shutdown: %s", err)
	}

	log.Info("server shutdown")
}

func initBusinessLogic(
	router *mux.Router,

	mid middleware.Middleware,
	httpResp response.HttpResponse,
	convert transperr.ErrorConverter,
	validationFormat strfmt.Registry,

	log logger.Logger,
) {
	webhookService := webhook.NewWebhookService()

	runner.InitHandlers(router, mid,
		webhook.NewRunnerHandlerV1(router, httpResp, convert, log, validationFormat, webhookService),
	)
}
