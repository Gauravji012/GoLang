package main

//ye go ke ander entry point hota h. yha se kisi bhi project ka execcution start ho jata h
//config file --> kisi bhi project ke liye configuration mgmt bhut jruri hota h kisi bhi project ke liye
//we use file based config system

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Gauravji012/GoLang/GoProjects/Student-api/internal/config"
	"github.com/Gauravji012/GoLang/GoProjects/Student-api/internal/http/handlers/student"
)

func main() {
	fmt.Println("Welcome to students api ...")
	/*
		step which we are follow in this project
		1. load config
		2. setup logger if needed -> we use inbuilt logger so dont required this step here
		3. database setup
		4. setup router
		5. setup server
	*/

	// load config
	cfg := config.MustLoad()

	// setup router
	// in go, net_http package is present which can be used for routing, creating servers.
	// before this version, we use 3rd party application to handle server side logic like routing, server client etc. query parameter / path parameters ko get krne ke liye bhut sara kaam krna pdta tha, if else use krna pdta tha.
	// lekin ab jo method usme they already provide methods to handle all of these stuff.

	//setting up router
	router := http.NewServeMux()

	// router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Welcome to students api..."))
	// }) -> this is not a good practice to create end point like this & write handler in this way. we should code like it should be maintainable, scalable.

	router.HandleFunc("POST /api/students", student.New())
	// setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	// fmt.Printf("server started successfully %s", cfg.HTTPServer.Addr) // suitable for simple, adhoc printing during development or for basic console output
	slog.Info("server started", slog.String("address", cfg.Addr)) // structured logging used in production

	done := make(chan os.Signal, 1) // to protect the program stop uninteruptly

	// getting response in done channel using notify method of signal
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done
	// here successfully complete the exeecution we start the process of server stopping

	slog.Info("shutting down the server") //gracefull shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // port is not block always with this approach if we get any issue while shutting down the server & it returns one function which we take in cancel variable
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}
	// also write like this in single line
	// err := server.Shutdown(ctx);err != nil {
	// 	slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	// }

	slog.Info("server shutdown successfully.. !!")
}
