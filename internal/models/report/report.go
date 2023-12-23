package report

import "time"

type Report struct {
	ID           int
	CreatedBy    int
	PostID       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ModeratorMsg string
	AdminMsg     string
	Status       string
}
