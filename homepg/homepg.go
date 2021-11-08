package homepg

import (
	"log"
	"net/http"
	"time"
)

const message = "Hello fellow gophers!"

type HomePage struct {
	Logger *log.Logger
}

func (h *HomePage) Handler(w http.ResponseWriter, r *http.Request) {
	begTime := time.Now()
	h.Logger.Printf("processing '/' request at %s\n", begTime.Format("YYYY-MM-dd hh:mm"))
	w.Header().Add("Content-Type", "text-plain; charset=utf-8")
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
	h.Logger.Printf("finished '/' in %v ms\n", time.Since(begTime))
}
