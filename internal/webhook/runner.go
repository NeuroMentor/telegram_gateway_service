package webhook

import (
	"telegram_gateway_service/internal/runner"
	"telegram_gateway_service/internal/shared/response"
	transperr "telegram_gateway_service/internal/shared/transport_error"
	"telegram_gateway_service/pkg/logger"

	"github.com/go-openapi/strfmt"
	"github.com/gorilla/mux"
)

type handlerV1 struct {
	router *mux.Router

	httpResp response.HttpResponse
	log      logger.Logger

	HttpHandler
}

func NewRunnerHandlerV1(
	router *mux.Router,

	httpResp response.HttpResponse,
	converter transperr.ErrorConverter,

	log logger.Logger,
	validationFormat strfmt.Registry,

	webhookService Service,
) runner.Handler {
	return &handlerV1{
		router: router.PathPrefix("/v1").Subrouter(),

		httpResp: httpResp,

		log: log,

		HttpHandler: NewWebhookHandler(httpResp, converter, webhookService, validationFormat),
	}
}

func (m *handlerV1) Init() []runner.Runner {
	return []runner.Runner{
		m.HttpHandler,
	}
}

func (m *handlerV1) RouterWithVersion() *mux.Router {
	return m.router
}
