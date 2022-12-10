package main

import "github.com/1saifj/go-swag-redoc/redoc"

func main() {
	//https://petstore.swagger.io/v2/swagger.json

	r := redoc.InitializeRedoc{
		Title:                 "Redoc",
		OpenAPISpeLocation:    "https://petstore.swagger.io/v2/swagger.json",
		RedocTemplateLocation: "./redoc.html",
	}
	r.CreateRedocTemplate()
}
