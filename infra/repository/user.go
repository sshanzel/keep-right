package repository

import (
	"github.com/sshanzel/keep-right/domain/entities"
	"github.com/sshanzel/keep-right/infra/db"
)

type UserRepository struct {
	context *db.MainContext
}

func NewUserRepository(ctx *db.MainContext) *UserRepository {
	ur := UserRepository{context: ctx}

	return &ur
}

func (_ur UserRepository) GetUsers() {

}
