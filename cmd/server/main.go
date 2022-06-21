package main

import (
	"fmt"
	"net/http"
)

type App struct {
}

func (app *App) Run() error {
	fmt.Println("App is running")
	handler := transportHttp.newHandler()
	handler.setUpRoutes()

	if err := http.ListenAndServe(":8000", handler.Router); err != nil {
		fmt.Println("Fail seet up server")
		return nil
	}
	return nil
}

func main() {
	fmt.Println("Test api")
	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("App not runing")
		fmt.Println(err)
	}

}
