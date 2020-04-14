package rest

import (
	"github.com/gorilla/mux"
	"github.com/jrollin/craft-challenge/adapters/rest/utils"
	"github.com/jrollin/craft-challenge/application/port_in"
	"log"
	"net/http"
)

type ListGamePlayersHandler struct {
	l  *log.Logger
	lg port_in.ListGamePlayers
	fg port_in.FindGame
}

func NewListGamePlayersHandler(log *log.Logger, lister port_in.ListGamePlayers, finder port_in.FindGame) *ListGamePlayersHandler {
	return &ListGamePlayersHandler{
		l:  log,
		lg: lister,
		fg : finder,
	}
}

func (gh *ListGamePlayersHandler) ListGamePlayers(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	code := vars["code"]
	gh.l.Printf("[DEBUG] list players fo game %s", code)

	g, err := gh.fg.FindByCode(code)
	if err != nil {
		gh.l.Printf("[ERROR] Game not found  %s", err)
		http.Error(rw, "Error finding game", http.StatusBadRequest)
		return
	}

	players, err := gh.lg.ListGamePlayers(g)
	if err != nil {
		gh.l.Printf("[ERROR] listing game players %s", err)
		http.Error(rw, "Error listing games", http.StatusBadRequest)
		return
	}

	//// map domain to api representation
	//gamesApi := api.NewGameListFromDomain(games)

	err = utils.ToJSON(players, rw)
	if err != nil {
		// we should never be here but log the error just incase
		gh.l.Printf("[ERROR] serializing game players %s", err)
		http.Error(rw, "Error listing games", http.StatusBadRequest)
	}
}
