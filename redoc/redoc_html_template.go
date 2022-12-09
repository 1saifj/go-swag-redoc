package redoc

import (
	"html/template"
	"os"
)

func ReadTemplate() {

	err := os.Chdir("./templates")
	if err != nil {
		return
	}

	tmpl, err := template.ParseFiles("redoc.html")
	if err != nil {
		panic(err)
	}
	// Define the data to pass to the template.
	data := map[string]interface{}{
		"Title": "My HTML Document",
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
