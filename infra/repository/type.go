package repository

import (
	"github.com/google/uuid"
	"github.com/sshanzel/keep-right/domain/entities"
	"github.com/sshanzel/keep-right/infra/db"
)

// ISomethingTypeRepository is the interface for SomethingTypeRepository implementation
type ISomethingTypeRepository interface {
	GetSomethingType(id uuid.UUID) *entities.SomethingType
	GetSomethingTypes() (somethingTypes []*entities.SomethingType)
	GetSomethingTypesOfUser(userID uuid.UUID) (somethingTypes []*entities.SomethingType)
	CreateSomethingType(somethingType entities.SomethingType) *entities.SomethingType
	UpdateSomethingType(somethingType entities.SomethingType)
	DeleteSomethingType(somethingType entities.SomethingType)
}

// SomethingTypeRepository is the handle for accessing DB Layer
type SomethingTypeRepository struct {
	ctx *db.MainContext
}

// NewSomethingTypeRepository creates a new instance of the Repository
func NewSomethingTypeRepository() *SomethingTypeRepository {
	sr := SomethingTypeRepository{ctx: db.Connect()}

	return &sr
}

// GetSomethingType fetches the user based on PK UUID
func (_sr SomethingTypeRepository) GetSomethingType(id uuid.UUID) *entities.SomethingType {
	somethingType := &entities.SomethingType{ID: id}
	err := _sr.ctx.DB.Select(somethingType)

	if err != nil {
		panic(err)
	}

	return somethingType
}

// GetSomethingTypes fetches all SomethingType
func (_sr SomethingTypeRepository) GetSomethingTypes() (somethingTypes []*entities.SomethingType) {
	err := _sr.ctx.DB.Model(&somethingTypes).Select()

	if err != nil {
		panic(err)
	}

	return
}

// GetSomethingTypesOfUser returns all SomethingType of the specified User
func (_sr SomethingTypeRepository) GetSomethingTypesOfUser(userID uuid.UUID) (somethingTypes []*entities.SomethingType) {
	somethingType := &entities.SomethingType{CreatedByID: userID}
	err := _sr.ctx.DB.Model(&somethingTypes).Select(somethingType)

	if err != nil {
		panic(err)
	}

	return
}

// CreateSomethingType insert data into DB
func (_sr SomethingTypeRepository) CreateSomethingType(somethingType *entities.SomethingType) {
	err := _sr.ctx.DB.Insert(somethingType)

	if err != nil {
		panic(err)
	}
}

// UpdateSomethingType updates the specified SomethingType
func (_sr SomethingTypeRepository) UpdateSomethingType(somethingType *entities.SomethingType) {
	err := _sr.ctx.DB.Update(somethingType)

	if err != nil {
		panic(err)
	}
}

// DeleteSomethingType removes the data from DB
func (_sr SomethingTypeRepository) DeleteSomethingType(somethingType *entities.SomethingType) {
	err := _sr.ctx.DB.Delete(somethingType)

	if err != nil {
		panic(err)
	}
}
