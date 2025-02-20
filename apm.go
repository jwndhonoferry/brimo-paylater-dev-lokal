package main

import (
	"log"
	"net/url"
	"os"

	"go.elastic.co/apm"
	"go.elastic.co/apm/transport"
)

func InitApm() {
	// Reload Environment variables to Elastic Apm.
	if _, err := transport.InitDefault(); err != nil {
		log.Fatal(err)
	}
	// Set ServiceName and ServiceVersion vars. from .env file then Run new Apm Tracer.
	serviceName := os.Getenv("ELASTIC_APM_SERVICE_NAME")
	serviceVersion := os.Getenv("ELASTIC_APM_SERVICE_VERSION")
	tracer, err := apm.NewTracer(serviceName, serviceVersion)
	if err != nil {
		panic(err)
	}
	transport, err := transport.NewHTTPTransport()
	if err != nil {
		panic(err)
	}

	//transport.SetSecretToken(os.Getenv("ELASTIC_APM_SECRET_TOKEN"))
	u, err := url.Parse(os.Getenv("ELASTIC_APM_SERVER_URL"))
	if err != nil {
		panic(err)
	}

	transport.SetServerURL(u)
	tracer.Transport = transport
	apm.DefaultTracer = tracer
	apm.DefaultTracer.SetCaptureBody(apm.CaptureBodyAll)
}
