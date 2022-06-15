package main

import "fmt"

type App struct {
}

func (app *App) Run() error {
	fmt.Println("App is running")
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
