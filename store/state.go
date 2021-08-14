package store

import (
	"main/models"
)

var Store State

type State struct {
	User models.User
}
