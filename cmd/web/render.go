package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type templateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	IsAuthenticated int
	API             string
	CSSVersion      string
}

var functions = template.FuncMap{}

//go:embed templates
var templateFS embed.FS

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	return td
}

func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, page string, td *templateData, partials ...string) error {
	var t *template.Template
	var err error
	templateToRender := fmt.Sprintf("templates/%s.page.tmpl", page)

	// I'll ignore the first return parameter,
	// I will get bool value of existing template page
	_, templateInMap := app.templateCache[templateToRender]

	// in production I don't want to be always stopping and starting this application every time
	// I make a change to the file in the base layout. Instead I want it to happen automatically
	if app.config.env == "production" && templateInMap {
		t = app.templateCache[templateToRender]
	} else {
		t, err = app.parseTemplate(partials, page, templateToRender)
		if err != nil {
			app.errorLog.Println(err)
			return err
		}
	}

	// check if data was passed with the call to add default data
	// if it wasn't, we create an empty template data object
	if td == nil {
		td = &templateData{}
	}

	// add our default data
	td = app.addDefaultData(td, r)

	// after adding default data finally we can execute that template
	// and return error if exits
	err = t.Execute(w, td)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}

	return nil
}

func (app *application) parseTemplate(partials []string, page string, templateToRender string) (*template.Template, error) {
	var t *template.Template
	var err error

	//build partials
	if len(partials) > 0 {
		for i, x := range partials {
			partials[i] = fmt.Sprintf("templates/%s.partial.tmpl", x)
		}
	}

	if len(partials) > 0 {
		t, err = template.New(fmt.Sprintf("%s.page.tmpl", page)).Funcs(functions).ParseFS(templateFS, "templates/base.layout.tmpl", strings.Join(partials, ","), templateToRender)
	} else {
		t, err = template.New(fmt.Sprintf("%s.page.tmpl", page)).Funcs(functions).ParseFS(templateFS, "templates/base.layout.tmpl", templateToRender)
	}

	if err != nil {
		app.errorLog.Println(err)
		return nil, err
	}

	app.templateCache[templateToRender] = t
	return t, nil
}
