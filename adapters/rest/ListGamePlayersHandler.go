package rest

import (
	"github.com/jrollin/craft-challenge/domain/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jrollin/craft-challenge/adapters/rest/utils"
	"github.com/jrollin/craft-challenge/domain/port_in/query"
)

type ListGamePlayersHandler struct {
	l  *log.Logger
	lg query.ListGamePlayers
	fg query.FindGameByCode
}

func NewListGamePlayersHandler(log *log.Logger, lister query.ListGamePlayers, finder query.FindGameByCode) *ListGamePlayersHandler {
	return &ListGamePlayersHandler{
		l:  log,
		lg: lister,
		fg: finder,
	}
}

// swagger:route GET /games/{code}/players player game playerListId
//
// List all players for a game
//
// Responses:
// 	default: genericErrorResponse
//  200: playerListResponse
//  404: notFoundResponse
//  422: validationErrorResponse
func (gh *ListGamePlayersHandler) ListGamePlayers(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	code := vars["code"]
	gh.l.Printf("[DEBUG] list players fo game %s", code)

	g, err := gh.fg.FindGameByCode(model.GameCode(code))
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
