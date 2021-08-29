package email

import (
	"io"
	"log"
	"net/http"
	"time"
)

type Handlers struct {
	logger *log.Logger
}

func (h *Handlers) EmailHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw, `{"message:" "sending a email! ❤"}`)
}

func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next(rw, r)
		h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
	}
}

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Logger(h.EmailHandler))
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{logger}
}
