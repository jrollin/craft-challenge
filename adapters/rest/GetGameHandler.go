package rest

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jrollin/craft-challenge/adapters/rest/api"
	"github.com/jrollin/craft-challenge/adapters/rest/utils"
	"github.com/jrollin/craft-challenge/application/port_in/query"
	"github.com/jrollin/craft-challenge/domain"
)

type GetGameHandler struct {
	l *log.Logger
	f query.FindGame
}

func NewGetGameHandler(log *log.Logger, finder query.FindGame) *GetGameHandler {
	return &GetGameHandler{
		l: log,
		f: finder,
	}
}

// swagger:route GET /games/{code} game gameId

// Get game by its code

// Responses:
//	200: gameResponse
//	404: notFoundResponse

// ListSingle handles GET requests
func (gh *GetGameHandler) GetGameByCode(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	code := vars["code"]
	gh.l.Printf("[DEBUG] get game by code %s", code)

	// find game by code received
	game, err := gh.f.FindByCode(domain.GameCode(code))
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
