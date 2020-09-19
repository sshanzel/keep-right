package services

import (
	"os"
)

var _config = "config"
var _dbname = "hGxqzZMWOU5YNv17QZNF"
var _dbaddr = "bm1ZvG2SscCO8f9M10s9"
var _dbuser = "AalYkK89XETcsLQuhtjd"
var _dbpassword = "Vgx74HLMB63pNxxH9aph"

// GetDbName returns GetDbName
func GetDbName() string {
	if os.Getenv("GAE_ENV") == "" {
		return ""
	}

	app, err := GetFirebaseApp()
	if err != nil {
		panic(err)
	}

	client, err := app.Firestore(_ctx)
	if err != nil {
		panic(err)
	}

	doc, err := client.Collection(_config).Doc(_dbname).Get(_ctx)
	if err != nil {
		return ""
	}

	data := doc.Data()
	if data == nil {
		return ""
	}

	return data["value"].(string)
}

// GetDbAddr returns GetDbAddr
func GetDbAddr() string {
	if os.Getenv("GAE_ENV") == "" {
		return ""
	}

	app, err := GetFirebaseApp()
	if err != nil {
		panic(err)
	}

	client, err := app.Firestore(_ctx)
	if err != nil {
		panic(err)
	}

	doc, err := client.Collection(_config).Doc(_dbaddr).Get(_ctx)
	if err != nil {
		return ""
	}

	data := doc.Data()
	if data == nil {
		return ""
	}

	return data["value"].(string)
}

// GetDbUsername returns GetDbUsername
func GetDbUsername() string {
	if os.Getenv("GAE_ENV") == "" {
		return ""
	}

	app, err := GetFirebaseApp()
	if err != nil {
		panic(err)
	}

	client, err := app.Firestore(_ctx)
	if err != nil {
		panic(err)
	}

	doc, err := client.Collection(_config).Doc(_dbuser).Get(_ctx)
	if err != nil {
		return ""
	}

	data := doc.Data()
	if data == nil {
		return ""
	}

	return data["value"].(string)
}

// GetDbPassword returns GetDbPassword
func GetDbPassword() string {
	if os.Getenv("GAE_ENV") == "" {
		return ""
	}

	app, err := GetFirebaseApp()
	if err != nil {
		panic(err)
	}

	client, err := app.Firestore(_ctx)
	if err != nil {
		panic(err)
	}

	doc, err := client.Collection(_config).Doc(_dbpassword).Get(_ctx)
	if err != nil {
		return ""
	}

	data := doc.Data()
	if data == nil {
		return ""
	}

	return data["value"].(string)
}
