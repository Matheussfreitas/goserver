package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Password string
	Token    string	
	Active   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
