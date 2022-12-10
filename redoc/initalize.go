package redoc

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type InitializeRedoc struct {
	Title                 string
	OpenAPISpeLocation    string
	RedocTemplateLocation *string
	isUsingSwagger        *bool
}

func (r *InitializeRedoc) CreateRedocTemplate() {
	// Set a default value for the isUsingSwagger field
	if r.isUsingSwagger == nil {
		r.isUsingSwagger = new(bool)
		*r.isUsingSwagger = true
	}
	checkRedocLocation(r)
	tmpl := `
    <!DOCTYPE html>
<html>
  <head>
    <title>Redoc</title>
  
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <link
      href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700"
      rel="stylesheet"
    />
    <style>
      body {
        margin: 0;
        padding: 0;
      }
    </style>
  </head>
  <body>
    <redoc spec-url={{.SpecLocation}}></redoc>
    <script src="https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js"></script>
  </body>
</html>
    `
	// Parse the HTML template
	t, err := template.New("redoc").Parse(tmpl)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := struct {
		Title        string
		SpecLocation string
	}{
		Title:        r.Title,
		SpecLocation: r.OpenAPISpeLocation,
	}

	// Create the file
	file, err := os.Create("redoc.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	err = t.Execute(file, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("HTML file created successfully")
}

func checkRedocLocation(r *InitializeRedoc) {
	if r.RedocTemplateLocation == nil {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		r.RedocTemplateLocation = new(string)
		*r.RedocTemplateLocation = wd
	}

}
