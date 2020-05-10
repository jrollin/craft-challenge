package rest

import (
	"github.com/jrollin/craft-challenge/domain/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jrollin/craft-challenge/adapters/rest/utils"
	"github.com/jrollin/craft-challenge/domain/port_in/query"
)

type GetGameCurrentStoryHandler struct {
	l  *log.Logger
	ds query.DisplayCurrentStory
	fg query.FindGameByCode
}

func NewGetGameCurrentStoryHandler(log *log.Logger, display query.DisplayCurrentStory, finder query.FindGameByCode) *GetGameCurrentStoryHandler {
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

	g, err := gh.fg.FindGameByCode(model.GameCode(code))
	if err != nil {
		gh.l.Printf("[ERROR] Game not found  %s", err)
		http.Error(rw, "Error finding game", http.StatusBadRequest)
		return
	}

	//@todo get current connected player
	player := &model.Player{}

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
