package domain

import "time"

type Game struct {
	Id        int 		`json:"id"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
}


