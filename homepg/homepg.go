package homepg

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type person struct {
	First  string    `json:"first"`
	Last   string    `json:"last"`
	Middle string    `json:"middle"`
	Phone  string    `json:"phone"`
	Dob    time.Time `json:"dob"`
}

var p = person{
	First:  "Eliot",
	Last:   "Easterling",
	Middle: "D",
	Phone:  "234-703-9147",
	Dob:    time.Date(1982, 10, 4, 11, 30, 0, 0, time.FixedZone("EST", (4*60*60))),
}

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

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(p)
	if err != nil {
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
