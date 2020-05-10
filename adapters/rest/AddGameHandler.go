package rest

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/jrollin/craft-challenge/adapters/rest/utils"
	"github.com/jrollin/craft-challenge/domain/port_in/command"
)

type AddGameHandler struct {
	l *log.Logger
	s command.AddGame
}

func NewAddGameHandler(log *log.Logger, adder command.AddGame) *AddGameHandler {
	return &AddGameHandler{
		l: log,
		s: adder,
	}
}

// swagger:route POST /games game addGameId
//
// Add a new game
//
// Responses:
// 	default: genericErrorResponse
//  201: noContentResponse
//  422: validationErrorResponse
func (gh *AddGameHandler) AddGame(rw http.ResponseWriter, r *http.Request) {

	gh.l.Printf("[DEBUG] add new game %s", r.Method)

	// decode request with anonymous struct
	t := &AddGameRequest{}

	err := utils.FromJSON(t, r.Body)
	if err != nil {
		gh.l.Printf("[ERROR] error decoding request %s", err)
		http.Error(rw, "Error processing request", http.StatusUnprocessableEntity)
		return
	}

	// create command from request
	id := uuid.New()
	ag, err := command.NewAddGameCommand(id, t.Code)
	if err != nil {
		gh.l.Printf("[ERROR] error creating command %s", err)
		http.Error(rw, "Invalid data provided", http.StatusUnprocessableEntity)
		return
	}

	// call usecase with valid command
	err = gh.s.AddGame(ag)
	if err != nil {
		gh.l.Printf("[ERROR] adding game failed %s", err)
		http.Error(rw, "Error adding game", http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

// An AddGameRequest model.
//
// swagger:parameters addGameId
type AddGameRequest struct {
	// The code to submit
	//
	// required: true
	// in: body
	Code string `json:"code"`
}
