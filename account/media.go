package account

import (
	"github.com/satori/go.uuid"
)

// Client account response struct
type Client struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}
