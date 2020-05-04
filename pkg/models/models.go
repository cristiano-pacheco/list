package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")
var IsNotPossibleToDelete = errors.New("models: is not possible to delete the resource")
var IsNotPossibleToCreate = errors.New("models: is not possible to create the resource")
var IsNotPossibleToUpdate = errors.New("models: is not possible to update the resource")

type List struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
