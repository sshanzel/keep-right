package repository

import (
	"github.com/google/uuid"
	"github.com/sshanzel/keep-right/domain/entities"
	"github.com/sshanzel/keep-right/infra/db"
)

// ISomethingRepository is the interface for SomethingRepository implementation
type ISomethingRepository interface {
	GetSomething(id uuid.UUID) *entities.Something
	GetSomethings() (somethings []*entities.Something)
	GetSomethingsOfUser(userID uuid.UUID) (somethings []*entities.Something)
	CreateSomething(something *entities.Something) *entities.Something
	UpdateSomething(something *entities.Something)
	DeleteSomething(something *entities.Something)
}

// SomethingRepository is the handle for accessing DB Layer
type SomethingRepository struct {
	ctx *db.MainContext
}

// NewSomethingRepository creates a new instance of the Repository
func NewSomethingRepository() *SomethingRepository {
	sr := SomethingRepository{ctx: db.Connect()}

	return &sr
}

// GetSomething fetches the user based on PK UUID
func (_sr SomethingRepository) GetSomething(id uuid.UUID) *entities.Something {
	something := &entities.Something{ID: id}
	err := _sr.ctx.DB.Select(something)

	if err != nil {
		panic(err)
	}

	return something
}

// GetSomethings fetches all Something
func (_sr SomethingRepository) GetSomethings() (somethings []*entities.Something) {
	err := _sr.ctx.DB.Model(&somethings).Select()

	if err != nil {
		panic(err)
	}

	return
}

// GetSomethingsOfUser returns all Something of the specified User
func (_sr SomethingRepository) GetSomethingsOfUser(userID uuid.UUID) (somethings []*entities.Something) {
	something := &entities.Something{CreatedByID: userID}
	err := _sr.ctx.DB.Model(&somethings).Select(something)

	if err != nil {
		panic(err)
	}

	return
}

// CreateSomething insert data into DB
func (_sr SomethingRepository) CreateSomething(something *entities.Something) *entities.Something {
	err := _sr.ctx.DB.Insert(something)

	if err != nil {
		panic(err)
	}

	return something
}

// UpdateSomething updates the specified Something
func (_sr SomethingRepository) UpdateSomething(something *entities.Something) {
	err := _sr.ctx.DB.Update(something)

	if err != nil {
		panic(err)
	}
}

// DeleteSomething removes the data from DB
func (_sr SomethingRepository) DeleteSomething(something *entities.Something) {
	err := _sr.ctx.DB.Delete(something)

	if err != nil {
		panic(err)
	}
}
