package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/niktheblak/temperature-api/internal/server"
	"github.com/niktheblak/temperature-api/pkg/measurement"
)

func main() {
	addr := os.Getenv("INFLUXDB_ADDR")
	if addr == "" {
		addr = "http://127.0.0.1:8086"
	}
	username := os.Getenv("INFLUXDB_USERNAME")
	password := os.Getenv("INFLUXDB_PASSWORD")
	database := os.Getenv("INFLUXDB_DATABASE")
	if database == "" {
		database = "ruuvitag"
	}
	meas := os.Getenv("INFLUXDB_MEASUREMENT")
	if meas == "" {
		meas = "ruuvitag"
	}
	port, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if port <= 0 || port > 65536 {
		port = 8080
	}
	cfg := measurement.Config{
		Addr:        addr,
		Username:    username,
		Password:    password,
		Database:    database,
		Measurement: meas,
		Timeout:     10 * time.Second,
	}
	svc, err := measurement.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	if err := svc.Ping(); err != nil {
		log.Fatal(err)
	}
	defer svc.Close()
	srv := &server.Server{
		Service: svc,
	}
	srv.Routes()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
