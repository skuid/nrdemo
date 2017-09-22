package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/newrelic/go-agent"
)

var appName = flag.String("app-name", "demoapp", "The New Relic app name")

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello!"}`))
	log.Printf("handled /hello: %s", r.RemoteAddr)
}

func main() {

	licenseKey := os.Getenv("NEW_RELIC_LICENSE_KEY")
	if licenseKey == "" {
		log.Fatalln("Missing environment variable NEW_RELIC_LICENSE_KEY!")
	}

	config := newrelic.NewConfig(*appName, licenseKey)
	app, err := newrelic.NewApplication(config)
	if err != nil {
		log.Fatalf("Error starting NR App: %s", err.Error())
	}

	http.HandleFunc(newrelic.WrapHandleFunc(app, "/", helloHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
