package handler

import (
	"net/http"
	"strings"
)

// ServeHTTPV2 is the HTTP router for all v2 api requests.
func (h *Handler) ServeHTTPV2(rw http.ResponseWriter, request *http.Request) {
	switch {
	case strings.HasPrefix(request.URL.Path, "/v2/agents"):
		http.StripPrefix("/v2", h.agentHandler).ServeHTTP(rw, request)
	case strings.HasPrefix(request.URL.Path, "/v2/host"):
		http.StripPrefix("/v2", h.hostHandler).ServeHTTP(rw, request)
	case strings.HasPrefix(request.URL.Path, "/v2/browse"):
		http.StripPrefix("/v2", h.browseHandler).ServeHTTP(rw, request)
	case strings.HasPrefix(request.URL.Path, "/v2/websocket"):
		http.StripPrefix("/v2", h.webSocketHandler).ServeHTTP(rw, request)
	}
}