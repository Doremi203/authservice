package types

import (
	"github.com/google/uuid"
)

type (
	UserID   uuid.UUID
	Email    string
	Password string
	AppID    int
	Token    string
)
