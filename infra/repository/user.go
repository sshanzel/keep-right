package repository

import (
	"github.com/google/uuid"
	"github.com/sshanzel/keep-right/domain/entities"
	"github.com/sshanzel/keep-right/infra/db"
)

// IUserRepository is the contract of implementation of UserRepository
type IUserRepository interface {
	GetUser(id uuid.UUID) *entities.User
	GetUsers() (users []*entities.User)
	CreateUser(user *entities.User) *entities.User
	UpdateUser(user *entities.User)
	DeleteUser(user *entities.User)
}

// UserRepository is the handle for User Queries
type UserRepository struct {
	ctx *db.MainContext
}

// NewUserRepository creates a new instance of the Repository
func NewUserRepository() *UserRepository {
	ur := UserRepository{ctx: db.Connect()}

	return &ur
}

// GetUser fetches all Users in the database
func (_ur UserRepository) GetUser(id uuid.UUID) *entities.User {
	user := &entities.User{ID: id}
	err := _ur.ctx.DB.Select(user)

	if err != nil {
		panic(err)
	}

	return user
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
