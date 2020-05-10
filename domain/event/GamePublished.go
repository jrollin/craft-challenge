package event

import (
	"github.com/jrollin/craft-challenge/domain/model"
	"time"
)

type GamePublished struct {
	OccuredAt time.Time
	GameId    model.GameID
}
