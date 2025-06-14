package webhook

import (
	"encoding/json"
	"net/http"
	"telegram_gateway_service/internal/runner"
	"telegram_gateway_service/internal/shared/dto"
	"telegram_gateway_service/internal/shared/middleware"
	"telegram_gateway_service/internal/shared/response"
	transperr "telegram_gateway_service/internal/shared/transport_error"
	"telegram_gateway_service/models"

	"github.com/go-openapi/strfmt"
	"github.com/gorilla/mux"
)

type HttpHandler interface {
	ProcessUpdate(w http.ResponseWriter, r *http.Request)
	ProcessMessage(w http.ResponseWriter, r *http.Request)
	ProcessCallback(w http.ResponseWriter, r *http.Request)
	ProcessEditedMessage(w http.ResponseWriter, r *http.Request)

	runner.Runner
}

type webhookHandler struct {
	httpResponse response.HttpResponse
	converter    transperr.ErrorConverter

	webhookService Service

	validationFormat strfmt.Registry
}

func NewWebhookHandler(
	httpResponse response.HttpResponse,
	converter transperr.ErrorConverter,

	webhookService Service,

	validationFormat strfmt.Registry,
) HttpHandler {
	return &webhookHandler{
		httpResponse: httpResponse,
		converter:    converter,

		webhookService: webhookService,

		validationFormat: validationFormat,
	}
}

func (m *webhookHandler) Run(router *mux.Router, mid middleware.Middleware) {

}

func (m *webhookHandler) ProcessUpdate(w http.ResponseWriter, r *http.Request) {
	var req models.TelegramUpdate

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		m.httpResponse.ErrorResponse(w, r,
			transperr.NewTransportError(transperr.ValidationError, http.StatusBadRequest),
		)
		return
	}

	err = req.Validate(m.validationFormat)
	if err != nil {
		m.httpResponse.ErrorResponse(w, r,
			transperr.NewTransportError(transperr.ValidationError, http.StatusBadRequest),
		)
		return
	}
	var update = dto.TelegramUpdateFromModels(&req)

	e := m.webhookService.ProcessUpdate(r.Context(), update)
	if e != nil {
		m.httpResponse.ErrorResponse(w, r, m.converter.ToHTTP(e))
		return
	}
	m.httpResponse.WriteResponse(w, r, http.StatusOK, nil)
}

func (m *webhookHandler) ProcessMessage(w http.ResponseWriter, r *http.Request) {
	var req models.TelegramMessage

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		m.httpResponse.ErrorResponse(w, r,
			transperr.NewTransportError(transperr.ValidationError, http.StatusBadRequest),
		)
		return
	}

	err = req.Validate(m.validationFormat)
	if err != nil {
		m.httpResponse.ErrorResponse(w, r,
			transperr.NewTransportError(transperr.ValidationError, http.StatusBadRequest),
		)
		return
	}
	var message = dto.TelegramMessageFromModels(&req)

	e := m.webhookService.ProcessMessage(r.Context(), message)
	if e != nil {
		m.httpResponse.ErrorResponse(w, r, m.converter.ToHTTP(e))
		return
	}
	m.httpResponse.WriteResponse(w, r, http.StatusOK, nil)
}

func (m *webhookHandler) ProcessCallback(w http.ResponseWriter, r *http.Request) {
	var req models.TelegramCallbackQuery

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		m.httpResponse.ErrorResponse(w, r,
			transperr.NewTransportError(transperr.ValidationError, http.StatusBadRequest),
		)
		return
	}

	err = req.Validate(m.validationFormat)
	if err != nil {
		m.httpResponse.ErrorResponse(w, r,
			transperr.NewTransportError(transperr.ValidationError, http.StatusBadRequest),
		)
		return
	}
	var callback = dto.TelegramCallbackQueryFromModels(&req)

	e := m.webhookService.ProcessCallback(r.Context(), callback)
	if e != nil {
		m.httpResponse.ErrorResponse(w, r, m.converter.ToHTTP(e))
		return
	}
	m.httpResponse.WriteResponse(w, r, http.StatusOK, nil)
}

func (m *webhookHandler) ProcessEditedMessage(w http.ResponseWriter, r *http.Request) {
	var req models.TelegramMessage

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		m.httpResponse.ErrorResponse(w, r,
			transperr.NewTransportError(transperr.ValidationError, http.StatusBadRequest),
		)
		return
	}

	err = req.Validate(m.validationFormat)
	if err != nil {
		m.httpResponse.ErrorResponse(w, r,
			transperr.NewTransportError(transperr.ValidationError, http.StatusBadRequest),
		)
		return
	}
	var message = dto.TelegramMessageFromModels(&req)

	e := m.webhookService.ProcessEditedMessage(r.Context(), message)
	if e != nil {
		m.httpResponse.ErrorResponse(w, r, m.converter.ToHTTP(e))
		return
	}
	m.httpResponse.WriteResponse(w, r, http.StatusOK, nil)
}
