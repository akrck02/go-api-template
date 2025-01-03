package models

import "github.com/akrck02/godot-api-template/constants"

type Error struct {
	Status  int                `json:"status,omitempty"`
	Error   constants.ApiError `json:"error,omitempty"`
	Message string             `json:"message,omitempty"`
}
