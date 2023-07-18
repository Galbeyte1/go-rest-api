package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Galbeyte1/go-rest-api/internal/comment"
	"github.com/Galbeyte1/go-rest-api/internal/database"
	transportHTTP "github.com/Galbeyte1/go-rest-api/internal/transport/http"
)

type App struct {}


func (app *App) Run() error {
	fmt.Println("Starting up Application")

	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Welcome to Go REST API")
	app := App{}
	if err := app.Run(); err != nil {
		log.Println("Error Starting up REST API")
		log.Println(err)
	}
}