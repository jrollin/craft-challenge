package rest

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jrollin/craft-challenge/application/port_in/command"
)

type PublishGameHandler struct {
	l *log.Logger
	c command.PublishGame
}

func NewPublishGameHandler(log *log.Logger, cmd command.PublishGame) *PublishGameHandler {
	return &PublishGameHandler{
		l: log,
		c: cmd,
	}
}

// swagger:route POST /games/{id}/publish game publishGameId
//
// Publish a new game
//
// Responses:
// 	default: genericErrorResponse
//  201: noContentResponse
//  422: validationErrorResponse
func (gh *PublishGameHandler) PublishGame(rw http.ResponseWriter, r *http.Request) {

	gh.l.Printf("[DEBUG] publish a game %s", r.Method)

	vars := mux.Vars(r)
	id := vars["id"]

	c, err := command.NewPublishGameCommand(uuid.MustParse(id), time.Now())
	if err != nil {
		gh.l.Printf("[ERROR] error command request %s", err)
		http.Error(rw, "Error processing command", http.StatusUnprocessableEntity)
		return
	}

	// call usecase with valid command
	err = gh.c.PublishGame(c)
	if err != nil {
		gh.l.Printf("[ERROR] Publish game failed %s", err)
		http.Error(rw, "Error publishing game", http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
