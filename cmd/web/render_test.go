package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddDefaultData(t *testing.T) {
	cfg := config{
		api: "https://example.com/api",
		stripe: struct {
			secret string
			key    string
		}{
			secret: "your_secret_key",
			key:    "your_publishable_key",
		},
	}

	var app = &application{
		config: cfg,
	}

	// Create an empty templateData struct
	td := &templateData{}

	// Create a mock HTTP request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Call the addDefaultData method
	result := app.addDefaultData(td, req)

	// Check if the API URL, Stripe secret key, and Stripe publishable key are set correctly
	assert.Equal(t, cfg.api, result.API)
	assert.Equal(t, cfg.stripe.secret, result.StripeSecretKey)
	assert.Equal(t, cfg.stripe.key, result.StripePublishableKey)
}

func TestParseTemplate(t *testing.T) {
	// Initialize the application
	app := &application{
		templateCache: make(map[string]*template.Template),
		errorLog:      log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	// Mock partials and templateToRender
	partials := []string{"stripe-js"}
	page := "terminal"
	templateToRender := "templates/terminal.page.gohtml"

	// Call the parseTemplate function
	tmpl, err := app.parseTemplate(partials, page, templateToRender)

	// Check if the returned template is not nil
	assert.NotNil(t, tmpl)

	// Check if the error is nil
	assert.NoError(t, err)

	// Check if the templates is cached
	cachedTemplate, ok := app.templateCache[templateToRender]
	assert.True(t, ok)
	assert.Equal(t, tmpl, cachedTemplate)

	// Test parsing with an invalid template
	invalidTemplate, err := app.parseTemplate(partials, "invalidPage", templateToRender)
	assert.Nil(t, invalidTemplate)
	assert.Error(t, err)
}
