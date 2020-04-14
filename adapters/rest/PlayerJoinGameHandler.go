package rest

import (
	"github.com/google/uuid"
	"github.com/jrollin/craft-challenge/adapters/rest/utils"
	"github.com/jrollin/craft-challenge/application/port_in"
	"log"
	"net/http"
)

type PlayerJoinGameHandler struct {
	l  *log.Logger
	pj port_in.JoinPlayerToGame
}

func NewPlayerJoinGameHandler(l *log.Logger, pj port_in.JoinPlayerToGame) *PlayerJoinGameHandler {
	return &PlayerJoinGameHandler{l: l, pj: pj}
}

// swagger:route POST /games/{code}/players game player addPlayerId
//
// Add a new player to game
//
// Responses:
// 	default: genericErrorResponse
//  201: noContentResponse
//  404: notFoundResponse
//  422: validationErrorResponse
func (gh *PlayerJoinGameHandler) JoinPlayerGame(rw http.ResponseWriter, r *http.Request) {

	gh.l.Printf("[DEBUG] player joins game %s", r.Method)

	// decode request with anonymous struct
	t := struct {
		Code     string `json:"code"`
		Username string `json:"username"`
		Server   string `json:"server"`
	}{}
	err := utils.FromJSON(&t, r.Body)
	if err != nil {
		gh.l.Printf("[ERROR] error decoding request : %s", err)
		http.Error(rw, "Error processing request", http.StatusUnprocessableEntity)
		return
	}

	// create command from request
	id := uuid.New()
	jg, err := port_in.NewJoinGameCommand(id, t.Code, t.Username, t.Server)
	if err != nil {
		gh.l.Printf("[ERROR] error creating command  %s", err)
		http.Error(rw, "Invalid data provided", http.StatusUnprocessableEntity)
		return
	}

	// call usecase with valid command
	err = gh.pj.JoinPlayerToGame(jg)
	if err != nil {
		gh.l.Printf("[ERROR] Join player to game failed %s ", err)
		http.Error(rw, "Error joining game", http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}


// An AddGameRequest model.
//
// swagger:parameters addPlayerId
type JoinPlayerRequest struct {
	// required: true
	// in: body
	Code     string `json:"code"`
	// required: true
	// in: body
	Username string `json:"username"`
	// required: true
	// in: body
	Server   string `json:"server"`
}

