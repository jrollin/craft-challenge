package event

import (
	"time"

	"github.com/jrollin/craft-challenge/domain"
)

type GamePublished struct {
	occuredAt time.Time
	GameId    domain.GameID
}
