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
	log    *log.Logger
	finder port_in.FindGame
}

func NewGetGameHandler(log *log.Logger, finder port_in.FindGame) *GetGameHandler {
	return &GetGameHandler{
		log:    log,
		finder: finder,
	}
}

func (g *GetGameHandler) GetGameByCode(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	code := vars["code"]
	g.log.Printf("[DEBUG] get game %s", code)

	// find game by code received
	game, err := g.finder.Find(code)
	if err != nil {
		g.log.Println("[ERROR] Find game", err)
		http.Error(rw, "Error finding game", http.StatusNotFound)
		return
	}

	// map domain to api representation
	gameApi := api.NewGame(game)

	// serialize
	err = utils.ToJSON(gameApi, rw)
	if err != nil {
		// we should never be here but log the error just incase
		g.log.Println("[ERROR] serializing game", err)
		http.Error(rw, "Error finding game", http.StatusBadRequest)
	}

}
