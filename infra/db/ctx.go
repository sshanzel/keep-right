package db

import (
	"os"

	"github.com/sshanzel/keep-right/domain/entities"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

var database = os.Getenv("database")
var dburi = os.Getenv("dburi")
var connStr = os.Getenv("dbconnstr")

// MainContext handles the database context for main services
type MainContext struct {
	DB *pg.DB
}

var connection *MainContext

// Connect opens the connection on the database and returns a Context
func Connect() *MainContext {
	if connection == nil {
		db := pg.Connect(&pg.Options{
			User:     "postgres",
			Password: "dr0w$$Ap",
			Database: "keep-right",
		})

		connection = &MainContext{DB: db}
	}

	// err := createSchema(_db)
	// if err != nil {
	// 	panic(err)
	// }

	return connection
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*entities.User)(nil), (*entities.Something)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}
