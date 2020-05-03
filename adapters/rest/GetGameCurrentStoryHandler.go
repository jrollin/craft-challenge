package rest

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jrollin/craft-challenge/adapters/rest/utils"
	"github.com/jrollin/craft-challenge/application/port_in"
	"github.com/jrollin/craft-challenge/domain"
)

type GetGameCurrentStoryHandler struct {
	l  *log.Logger
	ds port_in.DisplayCurrentStory
	fg port_in.FindGame
}

func NewGetGameCurrentStoryHandler(log *log.Logger, display port_in.DisplayCurrentStory, finder port_in.FindGame) *GetGameCurrentStoryHandler {
	return &GetGameCurrentStoryHandler{
		l:  log,
		ds: display,
		fg: finder,
	}
}

// swagger:route GET /games/{code}/stories/current story game story playerStoryId
//
// get current game story for player
//
// Responses:
// 	default: genericErrorResponse
//  200: playerListResponse
//  404: notFoundResponse
//  422: validationErrorResponse
func (gh *GetGameCurrentStoryHandler) GetGameCurrentStory(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	code := vars["code"]
	gh.l.Printf("[DEBUG] get current story fo game %s", code)

	g, err := gh.fg.FindByCode(domain.GameCode(code))
	if err != nil {
		gh.l.Printf("[ERROR] Game not found  %s", err)
		http.Error(rw, "Error finding game", http.StatusBadRequest)
		return
	}

	//@todo get current connected player
	player := &domain.Player{}

	story, err := gh.ds.DisplayCurrentStoryForPlayer(g, player)
	if err != nil {
		gh.l.Printf("[ERROR] get current story %s", err)
		http.Error(rw, "Error getting current story", http.StatusBadRequest)
		return
	}

	// @todo
	//// map domain to api representation

	err = utils.ToJSON(story, rw)
	if err != nil {
		// we should never be here but log the error just incase
		gh.l.Printf("[ERROR] serializing current story %s", err)
		http.Error(rw, "Error getting story", http.StatusBadRequest)
	}
}
