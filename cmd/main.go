package main

import (
	"fmt"
	"github.com/1saifj/go-swag-redoc/redoc"
	"github.com/spf13/cobra"
)

func main() {
	//https://petstore.swagger.io/v2/sswagger.json

	var cmdGenerate = &cobra.Command{
		Use:   "generate",
		Short: "Generate redoc documentation from OpenAPI definitions or Swag",
		Long:  `Generate redoc documentation from OpenAPI definitions or Swag`,
		Run: func(cmd *cobra.Command, args []string) {
			// Code for generating redoc documentation goes here
			fmt.Println("Generating redoc documentation...")
		},
	}

	var rootCmd = &cobra.Command{Use: "redoc-go"}
	rootCmd.AddCommand(cmdGenerate)
	rootCmd.Execute()

	redoc.ReadTemplate()

}

func GenerateRedoc() {
}
