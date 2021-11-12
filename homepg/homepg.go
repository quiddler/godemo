package homepg

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/quiddler/godemo/person"
)

const (
	begFmt = "process '/' request from %s at %s\n"
	endFmt = "finished '/' in %d\n"
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
	w.Header().Add("Content-Type", "application/javascript; charset=utf-8")
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

	err := json.NewEncoder(w).Encode(person.New())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.Logger.Println("Json encoding failed in homepg")
	}
}

func (h *HomePage) Log(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		start := time.Now()

		h.Logger.Printf(begFmt, r.RemoteAddr, start.Format(time.RFC3339Nano))
		defer h.Logger.Printf(endFmt, time.Since(start))

		next(rw, r)
	}
}

func (h *HomePage) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Log(h.Handler))
}
