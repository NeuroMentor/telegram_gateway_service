package response

import (
	"encoding/json"
	"net/http"
	transperr "telegram_gateway_service/internal/shared/transport_error"
	"telegram_gateway_service/pkg/logger"
)

type httpResponse struct {
	log     logger.Logger
	isDebug bool
}

func (m *httpResponse) ErrorResponse(w http.ResponseWriter, r *http.Request, err transperr.TransportError) {
	m.WriteResponse(w, r, err.GetCode(), err)
}

func (m *httpResponse) WriteResponse(w http.ResponseWriter, r *http.Request, code int, resp any) {
	raw, err := json.Marshal(resp)
	if err != nil {
		m.log.Errorf("error marshalling response: %v", err)
		return
	}

	pretty, _ := json.MarshalIndent(resp, "", "  ")
	if m.isDebug || code >= 500 {
		if code >= 500 {
			m.log.Errorf("[%s] '%s' Response: %s", r.Method, r.URL.String(), pretty)
		} else {
			m.log.Debugf("[%s] '%s' Response: %s", r.Method, r.URL.String(), pretty)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_, _ = w.Write(raw)
}

func NewHTTPResponse(log logger.Logger, isDebag bool) HttpResponse {
	return &httpResponse{
		log:     log,
		isDebug: isDebag,
	}
}
