package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T) {
	var app application

	mux := app.routes()

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing test passed
	default:
		t.Error(fmt.Sprintf("Type is not *chi.Mux, type is %T", v))
	}
}
