package rest

import (
	"github.com/gorilla/mux"
	"github.com/jrollin/craft-challenge/adapters/rest/api"
	"github.com/jrollin/craft-challenge/adapters/rest/utils"
	"github.com/jrollin/craft-challenge/application/port_in"
	"log"
	"net/http"
)

type GetGameHandler struct {
	l *log.Logger
	f port_in.FindGame
}

func NewGetGameHandler(log *log.Logger, finder port_in.FindGame) *GetGameHandler {
	return &GetGameHandler{
		l: log,
		f: finder,
	}
}

func (gh *GetGameHandler) GetGameByCode(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	code := vars["code"]
	gh.l.Printf("[DEBUG] get game by code %s", code)

	// find game by code received
	game, err := gh.f.FindByCode(code)
	if err != nil {
		gh.l.Printf("[ERROR] Find game by code %s", err)
		http.Error(rw, "Error finding game", http.StatusNotFound)
		return
	}

	// map domain to api representation
	gameApi := api.NewGameFromDomain(game)

	// serialize
	err = utils.ToJSON(gameApi, rw)
	if err != nil {
		// we should never be here but log the error just incase
		gh.l.Printf("[ERROR] serializing game %s", err)
		http.Error(rw, "Error finding game", http.StatusBadRequest)
	}

}
