package response

import (
	"net/http"
	transperr "telegram_gateway_service/internal/shared/transport_error"
)

type HttpResponse interface {
	ErrorResponse(w http.ResponseWriter, r *http.Request, err transperr.TransportError)
	WriteResponse(w http.ResponseWriter, r *http.Request, code int, resp any)
}
