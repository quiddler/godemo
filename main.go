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
var (
	KeyFile     = "./certs/ecdsa.pem"
	CertFile    = "./certs/mycert.crt"
	ServiceAddr = ":8080"
)

func main() {

	home := homepg.HomePage{
		Logger: log.New(os.Stdout, "demo - ", log.LstdFlags|log.Lshortfile),
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", home.Handler)

	s := server.New(mux, ServiceAddr)

	err := s.ListenAndServeTLS(CertFile, KeyFile)
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
