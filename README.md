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

## Contribution 

## License
go-swag-redoc is released under the [MIT License](https://opensource.org/licenses/MIT).
