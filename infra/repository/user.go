package repository

import (
	"github.com/sshanzel/keep-right/domain/entities"
	"github.com/sshanzel/keep-right/infra/db"
)

// UserRepository is the handle for User Queries
type UserRepository struct {
	ctx *db.MainContext
}

// NewUserRepository creates a new instance of the Repository
func NewUserRepository(_ctx *db.MainContext) *UserRepository {
	ur := UserRepository{ctx: _ctx}

	return &ur
}

// GetUsers fetches all Users in the database
func (_ur UserRepository) GetUsers() (users []entities.User) {
	err := _ur.ctx.DB.Model(&users)

	if err != nil {
		panic(err)
	}

	return
}
