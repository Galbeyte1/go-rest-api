package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/Galbeyte1/go-rest-api/internal/transport/http"
)

type App struct {}


func (app *App) Run() error {
	fmt.Println("Starting up Application")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Welcome to Go REST API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error Starting up REST API")
		fmt.Println(err)
	}
}