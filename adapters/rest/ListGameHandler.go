package rest

import (
	"github.com/jrollin/craft-challenge/adapters/rest/api"
	"github.com/jrollin/craft-challenge/adapters/rest/utils"
	"github.com/jrollin/craft-challenge/application/port_in"
	"log"
	"net/http"
)

type ListGameHandler struct {
	l  *log.Logger
	lg port_in.ListGames
}

func NewListGameHandler(log *log.Logger, lister port_in.ListGames) *ListGameHandler {
	return &ListGameHandler{
		l:  log,
		lg: lister,
	}
}


// swagger:route GET /games games GameListId
//
// List all available games
//
// Responses:
// 	default: genericError
//  200: gameListResponse
//  422: validationError
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
