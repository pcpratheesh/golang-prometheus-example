package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pcpratheesh/golang-prometheus-example/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	router := mux.NewRouter()
	PromMiddleware := middleware.NewPrometheusMiddleware()

	router.Handle("/metrics", promhttp.Handler())
	router.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("pong"))
	})
	router.Use(PromMiddleware.Metrics)

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("server initiate running...")
	log.Fatal(srv.ListenAndServe())
}
