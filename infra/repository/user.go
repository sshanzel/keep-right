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
func NewUserRepository() *UserRepository {
	ur := UserRepository{ctx: db.Connect()}

	return &ur
}

// GetUsers fetches all Users in the database
func (_ur UserRepository) GetUsers() (users []*entities.User) {
	err := _ur.ctx.DB.Model(&users).Select()

	if err != nil {
		panic(err)
	}

	return
}

// CreateUser returns the user with its ID and creates an item in the DB
func (_ur UserRepository) CreateUser(user *entities.User) *entities.User {
	err := _ur.ctx.DB.Insert(user)

	if err != nil {
		panic(err)
	}

	return user
}

// UpdateUser saves the changes for the specified User
func (_ur UserRepository) UpdateUser(user *entities.User) {
	err := _ur.ctx.DB.Update(user)

	if err != nil {
		panic(err)
	}
}

// DeleteUser removes the specified User in DB
func (_ur UserRepository) DeleteUser(user *entities.User) {
	err := _ur.ctx.DB.Delete(user)

	if err != nil {
		panic(err)
	}
}
