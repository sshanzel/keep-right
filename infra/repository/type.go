package repository

import (
	"github.com/google/uuid"
	"github.com/sshanzel/keep-right/domain/entities"
	"github.com/sshanzel/keep-right/infra/db"
)

// ISomeTypeRepository is the interface for SomeTypeRepository implementation
type ISomeTypeRepository interface {
	GetSomeType(id uuid.UUID) *entities.SomeType
	GetSomeTypes() (someTypes []*entities.SomeType)
	GetSomeTypesOfUser(userID uuid.UUID) (someTypes []*entities.SomeType)
	CreateSomeType(someType entities.SomeType) *entities.SomeType
	UpdateSomeType(someType entities.SomeType)
	DeleteSomeType(someType entities.SomeType)
}

// SomeTypeRepository is the handle for accessing DB Layer
type SomeTypeRepository struct {
	ctx *db.MainContext
}

// NewSomeTypeRepository creates a new instance of the Repository
func NewSomeTypeRepository() *SomeTypeRepository {
	sr := SomeTypeRepository{ctx: db.Connect()}

	return &sr
}

// GetSomeType fetches the user based on PK UUID
func (_sr SomeTypeRepository) GetSomeType(id uuid.UUID) *entities.SomeType {
	someType := &entities.SomeType{ID: id}
	err := _sr.ctx.DB.Select(someType)

	if err != nil {
		panic(err)
	}

	return someType
}

// GetSomeTypes fetches all SomeType
func (_sr SomeTypeRepository) GetSomeTypes() (someTypes []*entities.SomeType) {
	err := _sr.ctx.DB.Model(&someTypes).Select()

	if err != nil {
		panic(err)
	}

	return
}

// GetSomeTypesOfUser returns all SomeType of the specified User
func (_sr SomeTypeRepository) GetSomeTypesOfUser(userID uuid.UUID) (someTypes []*entities.SomeType) {
	someType := &entities.SomeType{CreatedByID: userID}
	err := _sr.ctx.DB.Model(&someTypes).Select(someType)

	if err != nil {
		panic(err)
	}

	return
}

// CreateSomeType insert data into DB
func (_sr SomeTypeRepository) CreateSomeType(someType *entities.SomeType) {
	err := _sr.ctx.DB.Insert(someType)

	if err != nil {
		panic(err)
	}
}

// UpdateSomeType updates the specified SomeType
func (_sr SomeTypeRepository) UpdateSomeType(someType *entities.SomeType) {
	err := _sr.ctx.DB.Update(someType)

	if err != nil {
		panic(err)
	}
}

// DeleteSomeType removes the data from DB
func (_sr SomeTypeRepository) DeleteSomeType(someType *entities.SomeType) {
	err := _sr.ctx.DB.Delete(someType)

	if err != nil {
		panic(err)
	}
}
