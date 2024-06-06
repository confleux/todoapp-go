package entities

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Description string    `json:"description"`
	Id          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	Uid         string    `json:"uid"`
}
