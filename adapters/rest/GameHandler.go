package rest

import (
	"github.com/gorilla/mux"
	"github.com/jrollin/craft-challenge/application/port_in"
	"github.com/jrollin/craft-challenge/domain"
	"log"
	"net/http"
)

type GameHandler struct {
	log    *log.Logger
	lister port_in.ListGames
	finder port_in.FindGame
	storer port_in.StoreGame
}

func NewGameHandler(log *log.Logger, lister port_in.ListGames, finder port_in.FindGame, storer port_in.StoreGame) *GameHandler {
	return &GameHandler{
		log:    log,
		lister: lister,
		finder: finder,
		storer: storer,
	}
}

func (g *GameHandler) ListAll(rw http.ResponseWriter, r *http.Request) {
	g.log.Println("[DEBUG] get all games")

	rw.Header().Add("Content-Type", "application/json")

	games, err := g.lister.GetAllGames()
	if err != nil {
		g.log.Print("[ERROR] listing games", err)
		http.Error(rw, "Error listing games", http.StatusBadRequest)
		return
	}

	err = ToJSON(games, rw)
	if err != nil {
		// we should never be here but log the error just incase
		g.log.Println("[ERROR] serializing game", err)
		http.Error(rw, "Error listing games", http.StatusBadRequest)
	}
}

func (g *GameHandler) GetGameByCode(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	code := vars["code"]
	g.log.Printf("[DEBUG] get game %s", code)

	rw.Header().Add("Content-Type", "application/json")

	game, err := g.finder.Find(code)
	if err != nil {
		g.log.Println("[ERROR] Find game", err)
		http.Error(rw, "Error finding game", http.StatusNotFound)
		return
	}

	err = ToJSON(game, rw)
	if err != nil {
		// we should never be here but log the error just incase
		g.log.Println("[ERROR] serializing game", err)
		http.Error(rw, "Error finding game", http.StatusBadRequest)
	}
}

func (g *GameHandler) AddGame(rw http.ResponseWriter, r *http.Request) {

	g.log.Printf("[DEBUG] add new game %s", r.Method)

	rw.Header().Add("Content-Type", "application/json")

	game := &domain.Game{}
	err := FromJSON(game, r.Body)
	if err != nil {
		g.log.Printf("[ERROR] error unmarshal json %s", err)
		http.Error(rw, "Invalid data provided", http.StatusUnprocessableEntity)
		return
	}

	err = g.storer.Store(game)
	if err != nil {
		g.log.Println("[ERROR] Store game failed", err)
		http.Error(rw, "Error storing game", http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
