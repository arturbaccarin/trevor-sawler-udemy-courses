package main

import (
	"os"
	"testing"
	"webapp/pkg/repository/dbrepo"
)

var app application

// alway be executed before the tests run
func TestMain(m *testing.M) {
	pathToTemplates = "./../../templates/"

	app.Session = getSession()
	app.DB = &dbrepo.TestDBRepo{}

	os.Exit(m.Run())
}
