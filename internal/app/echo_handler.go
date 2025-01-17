package app

import (
	"io"
	"net/http"

	"go.uber.org/zap"
)

type EchoHandler struct {
	log *zap.Logger
}

func NewEchoHandler(log *zap.Logger) *EchoHandler {
	return &EchoHandler{log: log}
}

func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		h.log.Warn("Failed to handle request", zap.Error(err))
		// fmt.Fprintln(os.Stderr, "Failed to handle request:", err)
	}
}

func (h *EchoHandler) Pattern() string {
	return "/echo"
}
