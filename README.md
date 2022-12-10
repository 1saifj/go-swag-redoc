## redoc-go
**redoc-go** is a Go package that generates **Redoc documentation** from Swag documentation or OpenAPI defnitions.

## Features
- Generates a complete [Redoc](https://redocly.com/redoc/) HTML documentation file (redoc.html) from Swag documentation.
- Command-line interface for easily generating the Redocly documentation.

## Installation

To install the redoc-go module, run the following command:
```sh
go get github.com/1saifj/redoc-go
```


## Usage

To use the redoc-go module in your Go project, import it into your code and use the `InitializeRedoc` struct and the `CreateRedocTemplate` method to configure and create the Redoc HTML template. You can then use the `RedocMiddleware` function to create a middleware handler and register it with your web framework's router or application.


Here's an example of how you could use the InitializeRedoc struct and the CreateRedocTemplate method to create a Redoc HTML template, and the RedocMiddleware function to create a middleware handler and register it with the http package:


```go
package main

import (
    "net/http"
    "github.com/1saifj/redoc-go"
)

func main() {
    // Initialize the Redoc documentation
    redocConfiguration := redoc.InitializeRedoc{
        Title:              "Redoc",
        OpenAPISpeLocation: "https://petstore.swagger.io/v2/swagger.json",
    }
    redocConfiguration.CreateRedocTemplate()

    // Use the RedocMiddleware function to create a middleware handler
    redocHandler := RedocMiddleware(nextHandler, "/redoc")

    // Register the middleware handler with the http package
    http.Handle("/", redocHandler)
    http.ListenAndServe(":8080", nil)
}
```

## Contribution 

## License
go-swag-redoc is released under the [MIT License](https://opensource.org/licenses/MIT).
