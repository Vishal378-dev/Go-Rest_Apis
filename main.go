package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vishal/Rest_Apis/internal/configs"
)

func main() {
	cfg := configs.ShouldLoad()
	// load config
	// logger package

	// db setup
	// route setting
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Write([]byte("Hello"))
		} else {
			w.Write([]byte("wrong method"))
		}

	})
	// server
	slog.Info("server started", slog.String("Address -", cfg.HttpServer.Address))
	server := http.Server{
		Addr:    cfg.HttpServer.Address,
		Handler: r,
	}
	// to gracefully stop the server due to some interuption
	complete := make(chan os.Signal, 1)
	signal.Notify(complete, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Error while starting the server - %s\n", err.Error())
		}
	}()
	<-complete

	//log before shutting the server
	//slog is used for structured log
	slog.Info("Shutting Down the Server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown - ", slog.String("error ", err.Error()))
	}

	slog.Info("Server ShutDown Successfully")
}
