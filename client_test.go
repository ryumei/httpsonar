package main

import (
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestGetOK(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/").
		Reply(200).
		JSON(map[string]string{"message": "Hello World!"})

	// normal case
	result, err := ping("http://example.com/", 200)
	if err != nil {
		t.Fatalf("Failed to GET")
	}
	if !result {
		t.Fatalf("Invalid status code")
	}
}

func TestGetNotFound(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/notfound.html").
		Reply(404).
		JSON(map[string]string{"msg": "Hello World!"})

	// error case: not found
	result, err := ping("http://example.com/notfound.html", 404)
	if err != nil {
		t.Fatalf("Failed to GET")
	}
	if !result {
		t.Fatalf("Invalid status code")
	}
}
