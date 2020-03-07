package db

import (
	"os"

	"github.com/sshanzel/keep-right/domain/entities"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type connStr struct {
	Username string
	Password string
	Database string
}

func getConnStr() (conn *connStr) {
	_username := os.Getenv("db:username")
	if _username == "" {
		_username = "postgres"
	}

	_password := os.Getenv("db:password")
	if _password == "" {
		_password = "dr0w$$Ap"
	}

	_db := os.Getenv("db:database")
	if _db == "" {
		_db = "keep-right"
	}

	return &connStr{_username, _password, _db}
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
			User:     conn.Username,
			Password: conn.Password,
			Database: conn.Database,
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
