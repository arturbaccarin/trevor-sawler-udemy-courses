package main

import (
	"os"
	"testing"
)

var app application

// alway be executed before the tests run
func TestMain(m *testing.M) {
	app.Session = getSession()

	os.Exit(m.Run())
}
