package main

import (
	"log"
	"net/http"

	"flag"

	"github.com/golang/glog"
	"github.com/n3wscott/cmd-service-broker/cmd/cmd-broker/server"
)

func main() {

	flag.Parse()

	s := server.CreateServer()
	glog.Infof("Starting Broker, %s", "http://localhost:12345")
	log.Fatal(http.ListenAndServe(":12345", s.Router))
}
