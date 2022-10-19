package main

import (
	"io/ioutil"
	"log"
	"lotr"
	"os"
	"testing"
)

const ACCESS_CODE = "mRNnJRcQoHPw3duyL6Vp"

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	code := m.Run()
	os.Exit(code)
}

func TestUnauthorized(t *testing.T) {
	app := lotr.New("")

	_, err := app.ListCharacters()
	if err == nil || err.Error() != "Unauthorized" {
		t.Fatal("Should return Unauthorized when no access code")
	}
}
