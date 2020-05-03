package rest

import (
	"log"
	"net/http"

	"github.com/jrollin/craft-challenge/adapters/rest/api"
	"github.com/jrollin/craft-challenge/adapters/rest/utils"
	"github.com/jrollin/craft-challenge/application/port_in/query"
)

type ListGameHandler struct {
	l  *log.Logger
	lg query.ListGames
}

func NewListGameHandler(log *log.Logger, lister query.ListGames) *ListGameHandler {
	return &ListGameHandler{
		l:  log,
		lg: lister,
	}
}

// swagger:route GET /games game GameListId
//
// List all available games
//
// Responses:
// 	default: genericErrorResponse
//  200: gameListResponse
//  422: validationErrorResponse
func (gh *ListGameHandler) ListAll(rw http.ResponseWriter, r *http.Request) {
	gh.l.Println("[DEBUG] list all games")

	games, err := gh.lg.GetAllGames()
	if err != nil {
		gh.l.Printf("[ERROR] listing games %s", err)
		http.Error(rw, "Error listing games", http.StatusBadRequest)
		return
	}

	// map domain to api representation
	gamesApi := api.NewGameListFromDomain(games)

	err = utils.ToJSON(gamesApi, rw)
	if err != nil {
		// we should never be here but log the error just incase
		gh.l.Printf("[ERROR] serializing game %s", err)
		http.Error(rw, "Error listing games", http.StatusBadRequest)
	}
}
