package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/jrollin/craft-challenge/adapters/persistence"
	"github.com/jrollin/craft-challenge/adapters/rest"
	"github.com/jrollin/craft-challenge/application/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var bindAddress = ":3000"

func main() {

	l := log.New(os.Stdout, "craft-api ", log.LstdFlags)
	l.Printf("Starting server on address %s", bindAddress)

	// adapter
	gr := persistence.NewGameRepositoryInMemoryAdapter()
	// use cases
	gl := service.NewGameLister(gr)
	gf := service.NewGameFinder(gr)
	gs := service.NewGameStorer(gr)

	// Rest handlers
	gh := rest.NewGameHandler(l, gl, gf, gs)

	// new router
	sm := mux.NewRouter()

	// define GET routes
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/games", gh.ListAll)
	getRouter.HandleFunc("/games/{code:[a-z]+}", gh.GetGameByCode)
	// define POST routes
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/games", gh.AddGame)

	// configure http server
	s := http.Server{
		Addr:     bindAddress,
		Handler:  sm,
		ErrorLog: l,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// start http server
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server %s ", err)
			os.Exit(1)
		}
	}()

	// graceful interruption
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// blocked until signal emitted
	sig := <-c
	log.Println("Got signal ", sig)

	// graceful shutdown server, waiting 30s
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := s.Shutdown(ctx)
	if err != nil {
		l.Printf("Error shutdown %s", err)
	}

}
