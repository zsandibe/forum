package request

import "time"

type Request struct {
	ID        int
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    string
}
