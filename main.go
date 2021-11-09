package main

import (
	"log"
	"net/http"
	"os"

	"github.com/quiddler/godemo/homepg"
	"github.com/quiddler/godemo/server"
)

// run the following command to get the .pem and .crt files that tls relies on:
// openssl req -x509 -nodes -newkey ec -pkeyopt ec_paramgen_curve:secp384r1 -keyout ecdsa.pem -out mycert.crt -days 3650
const (
	KeyFile     = "./certs/ecdsa.pem"
	CertFile    = "./certs/mycert.crt"
	ServiceAddr = ":8080"
)

func main() {

	logger := log.New(os.Stdout, "demo - ", log.LstdFlags|log.Lshortfile)

	mux := http.NewServeMux()

	h := homepg.New(logger)
	h.RegisterRoutes(mux)

	s := server.New(mux, ServiceAddr)

	logger.Fatal(s.ListenAndServeTLS(CertFile, KeyFile))
}
