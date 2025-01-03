package models

import "github.com/akrck02/godot-api-template/errors"

type Error struct {
	Status  int             `json:"status,omitempty"`
	Error   errors.ApiError `json:"error,omitempty"`
	Message string          `json:"message,omitempty"`
}
