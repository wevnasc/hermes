package server

import (
	"log"
	"net/http"
	"time"
)

type Middlewares struct {
	Log *log.Logger
}

func (m *Middlewares) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next(rw, r)
		m.Log.Printf("request processed in %s\n", time.Now().Sub(startTime))
	}
}

func (m *Middlewares) Headers(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-type", "application/json")
		next(rw, r)
	}
}

func (m *Middlewares) Method(next http.HandlerFunc, allowed []string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		for _, method := range allowed {
			if method == r.Method {
				next(rw, r)
				return
			}
		}

		rw.WriteHeader(http.StatusMethodNotAllowed)
		return

	}
}

func (m *Middlewares) Resource(next http.HandlerFunc, allowed []string) http.HandlerFunc {
	return m.Logger(m.Headers(m.Method(func(rw http.ResponseWriter, r *http.Request) {
		next(rw, r)
	}, allowed)))
}
