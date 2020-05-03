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

	// use cases
	gl := service.NewGameLister(l, gr)
	gf := service.NewGameFinder(l, gr)
	gs := service.NewGameAdder(l, gr)
	gst := service.NewGameStarter(l, gr)
	pgj := service.NewPlayerGameJoiner(l, gr, gr)
	pgl := service.NewGamePlayerLister(l, gr)
	pgr := service.NewGameStoryReader(l, gr)
	// related handlers
	glh := rest.NewListGameHandler(l, gl)
	gfh := rest.NewGetGameHandler(l, gf)
	gsh := rest.NewAddGameHandler(l, gs)
	gsth := rest.NewStartGameHandler(l, gst, gf)
	pgjh := rest.NewPlayerJoinGameHandler(l, pgj)
	pglh := rest.NewListGamePlayersHandler(l, pgl, gf)
	pgrh := rest.NewGetGameCurrentStoryHandler(l, pgr, gf)

	// new router
	sm := mux.NewRouter()

	// define GET routes
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/admin/games/{code:[a-z]+}/players", pglh.ListGamePlayers)
	getRouter.HandleFunc("/admin/games", glh.ListAll)
	getRouter.HandleFunc("/games/{code:[a-z]+}", gfh.GetGameByCode)
	getRouter.HandleFunc("/games/{code:[a-z]+}/stories/current", pgrh.GetGameCurrentStory)

	// define POST routes
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/admin/games", gsh.AddGame)
	postRouter.HandleFunc("/admin/games/{code:[a-z]+}/start", gsth.StartGame)
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
