package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database = os.Getenv("database")
var dburi = os.Getenv("dburi")

func getConfig(ctx context.Context) (*mongo.Client, error) {
	return mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
}

// MainContext handles the database context for main services
type MainContext struct {
	Users      *mongo.Collection
	Somethings *mongo.Collection
}

// New returns an instantiated DBContext
func New() (dbctx *MainContext, err error) {
	_ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	cancel()

	client, err := getConfig(_ctx)

	if err != nil {
		return
	}

	dbctx.Users = client.Database(database).Collection("users")
	dbctx.Somethings = client.Database(database).Collection("somethings")

	return
}
