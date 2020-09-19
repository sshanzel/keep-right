package db

import (
	"github.com/sshanzel/keep-right/domain/entities"
	"github.com/sshanzel/keep-right/infra/services"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type connStr struct {
	Addr     string
	Username string
	Password string
	Database string
}

func getConnStr() (conn *connStr) {
	addr := services.GetDbAddr()
	if addr == "" {
		addr = ":5432"
	}

	username := services.GetDbUsername()
	if username == "" {
		username = "postgres"
	}

	password := services.GetDbPassword()
	if password == "" {
		password = "dr0w$$Ap"
	}

	db := services.GetDbName()
	if db == "" {
		db = "keep-right"
	}

	return &connStr{addr, username, password, db}
}

// MainContext handles the database context for main services
type MainContext struct {
	DB *pg.DB
}

var connection *MainContext

// Connect opens the connection on the database and returns a Context
func Connect() *MainContext {
	if connection == nil {
		conn := getConnStr()
		db := pg.Connect(&pg.Options{
			Addr:     conn.Addr,
			User:     conn.Username,
			Password: conn.Password,
			Database: conn.Database,
		})

		err := createSchema(db)
		if err != nil {
			// panic(err)
		}

		connection = &MainContext{DB: db}
	}

	return connection
}

func createSchema(db *pg.DB) error {
	// err := db.CreateTable(model, &orm.CreateTableOptions{})
	err := db.CreateTable(&entities.User{}, &orm.CreateTableOptions{})
	err = db.CreateTable(&entities.SomethingType{}, &orm.CreateTableOptions{})
	err = db.CreateTable(&entities.Something{}, &orm.CreateTableOptions{})
	if err != nil {
		return err
	}
	return nil
}
