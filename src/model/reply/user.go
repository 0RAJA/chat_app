package reply

import (
	"time"
)

type CreateUser struct {
	ID       int64     `json:"id"`
	Email    string    `json:"email"`
	CreateAt time.Time `json:"create_at"`
}

type ExistEmail struct {
	Exist bool `json:"exist"`
}
