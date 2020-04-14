package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"

	"github.com/jrollin/craft-challenge/adapters/persistence"
	"github.com/jrollin/craft-challenge/adapters/rest"
	"github.com/jrollin/craft-challenge/application/service"
)

var bindAddress = ":3000"

func main() {

	l := log.New(os.Stdout, "craft-api ", log.LstdFlags)
	l.Printf("Starting server on address %s", bindAddress)

	// adapter
	gr := persistence.NewGameRepositoryInMemoryAdapter()

	// use cases with related handlers
	gl := service.NewGameLister(l, gr)
	glh := rest.NewListGameHandler(l, gl)

	gf := service.NewGameFinder(l, gr)
	gfh := rest.NewGetGameHandler(l, gf)

	gs := service.NewGameAdder(l, gr)
	gsh := rest.NewAddGameHandler(l, gs)

	pgj := service.NewPlayerGameJoiner(l, gr, gr)
	pgjh := rest.NewPlayerJoinGameHandler(l, pgj)

	pgl := service.NewGamePlayerLister(l, gr)
	pglh := rest.NewListGamePlayersHandler(l, pgl, gf)

	// new router
	sm := mux.NewRouter()

	// define GET routes
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/games", glh.ListAll)
	getRouter.HandleFunc("/games/{code:[a-z]+}", gfh.GetGameByCode)
	getRouter.HandleFunc("/games/{code:[a-z]+}/players", pglh.ListGamePlayers)

	// define POST routes
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/games", gsh.AddGame)
	postRouter.HandleFunc("/games/{code:[a-z]+}/players", pgjh.JoinPlayerGame)

	// doc
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	sm.Handle("/docs", sh)
	sm.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// configure http server
	s := http.Server{
		Addr:     bindAddress,
		Handler:  sm,
		ErrorLog: l,
		// Good practice: enforce timeouts for servers
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
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
