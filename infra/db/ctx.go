package db

import (
	"os"

	"github.com/sshanzel/keep-right/domain/entities"
)

var database = os.Getenv("database")
var dburi = os.Getenv("dburi")

// MainContext handles the database context for main services
type MainContext struct {
	Users      []entities.User
	Somethings []entities.Something
}
