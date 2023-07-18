package main

import (
	"fmt"
)

type App struct {}


func (app *App) Run() error {
	fmt.Println("Starting up Application")
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