package main

import (
	"flag"
	"fmt"
	"github.com/matfire/spatial/server"
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var configPath string
	var port int
	var debug bool
	var logFile string

	flag.StringVar(&configPath, "config", "./config.toml", `Path to config file (defaults to current dir's config.toml)`)
	flag.IntVar(&port, "port", 8080, "port to run the webserver on (defaults to 8080)")
	flag.BoolVar(&debug, "debug", false, "turns on debug mode for the web server (defaults to off)")
	flag.StringVar(&logFile, "log", "stdout", "where to write logs (defaults to stdout)")
	flag.Parse()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	if logFile != "stdout" {
		file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		logger = slog.New(slog.NewJSONHandler(file, nil))
		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}
	logger.Info("parsed flags")
	instance := server.NewServer(logger)
	logger.Info("starting server on port " + strconv.Itoa(port))
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), instance)
	if err != nil {
		panic("could not start server")
	}
}
