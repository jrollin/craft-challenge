package rest

import (
	"github.com/jrollin/craft-challenge/adapters/rest/api"
	"github.com/jrollin/craft-challenge/adapters/rest/utils"
	"github.com/jrollin/craft-challenge/application/port_in"
	"log"
	"net/http"
)

type ListGameHandler struct {
	log    *log.Logger
	lister port_in.ListGames
}

func NewListGameHandler(log *log.Logger, lister port_in.ListGames) *ListGameHandler {
	return &ListGameHandler{
		log:    log,
		lister: lister,
	}
}

func (g *ListGameHandler) ListAll(rw http.ResponseWriter, r *http.Request) {
	g.log.Println("[DEBUG] get all games")

	rw.Header().Add("Content-Type", "application/json")

	games, err := g.lister.GetAllGames()
	if err != nil {
		g.log.Print("[ERROR] listing games", err)
		http.Error(rw, "Error listing games", http.StatusBadRequest)
		return
	}

	// map domain to api representation
	gamesApi := api.NewGameList(games)

	err = utils.ToJSON(gamesApi, rw)
	if err != nil {
		// we should never be here but log the error just incase
		g.log.Println("[ERROR] serializing game", err)
		http.Error(rw, "Error listing games", http.StatusBadRequest)
	}
}
