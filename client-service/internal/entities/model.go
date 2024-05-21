package entities

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Description string
	Id          uuid.UUID
	CreatedAt   time.Time
	Uid         string
}
