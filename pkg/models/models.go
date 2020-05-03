package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type List struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
