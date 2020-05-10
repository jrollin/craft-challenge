package main

import (
	"context"
	"github.com/jrollin/craft-challenge/application"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"

	"github.com/jrollin/craft-challenge/adapters/persistence"
	"github.com/jrollin/craft-challenge/adapters/rest"
)

var bindAddress = ":3000"

func main() {

	l := log.New(os.Stdout, "craft-api ", log.LstdFlags)
	l.Printf("Starting server on address %s", bindAddress)

	// new router
	sm := mux.NewRouter()

	// adapter
	gr := persistence.NewGameRepositoryInMemoryAdapter()

	// query use cases
	ql := application.NewGameLister(l, gr)
	qf := application.NewGameFinder(l, gr)
	qpl := application.NewGamePlayerLister(l, gr)
	qsr := application.NewGameStoryReader(l, gr)
	// query handlers
	qlh := rest.NewListGameHandler(l, ql)
	qfh := rest.NewGetGameHandler(l, qf)
	qplh := rest.NewListGamePlayersHandler(l, qpl, qf)
	qsrh := rest.NewGetGameCurrentStoryHandler(l, qsr, qf)

	// define GET routes
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/games", qlh.ListAll)
	getRouter.HandleFunc("/games/{id}", qfh.GetGameByCode)
	getRouter.HandleFunc("/games/{id}/players", qplh.ListGamePlayers)
	getRouter.HandleFunc("/games/{id}/stories/current", qsrh.GetGameCurrentStory)

	// command use cases
	ca := application.NewGameAdder(l, gr)
	cs := application.NewGameStarter(l, gr, gr)
	cp := application.NewGamePublisher(l, gr, gr)
	cpj := application.NewPlayerGameJoiner(l, gr, gr)
	// command handlers
	cah := rest.NewAddGameHandler(l, ca)
	cph := rest.NewPublishGameHandler(l, cp)
	csh := rest.NewStartGameHandler(l, cs, qf)
	cpjh := rest.NewPlayerJoinGameHandler(l, cpj)

	// define POST routes
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/games", cah.AddGame)
	postRouter.HandleFunc("/games/{id}/publish", cph.PublishGame)
	postRouter.HandleFunc("/games/{id}/start", csh.StartGame)
	postRouter.HandleFunc("/games/{id}/players", cpjh.JoinPlayerGame)

	// doc
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	// redoc and swagger routes
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
