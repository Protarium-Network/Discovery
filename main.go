package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Endpoint struct {
	Host       string `xml:"host"`
	ApiHost    string `xml:"api_host"`
	PortalHost string `xml:"portal_host"`
	N3DSHost   string `xml:"n3ds_host"`
}

type Result struct {
	XMLName  xml.Name `xml:"result"`
	HasError int      `xml:"has_error"`
	Version  int      `xml:"version"`
	Endpoint Endpoint `xml:"endpoint"`
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func handleConfig(w http.ResponseWriter, r *http.Request) {
	apiHost := getEnv("API_HOST", "")
	portalHost := getEnv("PORTAL_HOST", "")
	n3dsHost := getEnv("N3DS_HOST", "")

	result := Result{
		HasError: 0,
		Version:  1,
		Endpoint: Endpoint{
			Host:       apiHost,
			ApiHost:    apiHost,
			PortalHost: portalHost,
			N3DSHost:   n3dsHost,
		},
	}

	output, err := xml.MarshalIndent(result, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Header().Set("X-Dispatch", "Olive::Web::Discovery::V1::Endpoint-index")
	w.Header().Set("Connection", "close")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(xml.Header))
	w.Write(output)
}

func main() {
	httpPort := getEnv("HTTP_PORT", "3002")

	http.HandleFunc("/", handleConfig)

	addr := fmt.Sprintf(":%s", httpPort)
	log.Printf("Server started on port %s...", httpPort)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
