package homepg

import (
	"log"
	"net/http"
	"time"
)

const (
	msg = "Hello fellow gophers!"
	beg = "process '/' request from %s at %s\n"
	end = "finished '/' in %d\n"
)

type HomePage struct {
	Logger *log.Logger
}

func New(logger *log.Logger) *HomePage {
	return &HomePage{
		Logger: logger,
	}
}

func (h *HomePage) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text-plain; charset=utf-8")
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

	w.WriteHeader(http.StatusOK)

	w.Write([]byte(msg))
}

func (h *HomePage) Log(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		start := time.Now()

		h.Logger.Printf(beg, r.RemoteAddr, start.Format(time.RFC3339Nano))
		defer h.Logger.Printf(end, time.Since(start))

		next(rw, r)
	}
}

func (h *HomePage) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Log(h.Handler))
}
