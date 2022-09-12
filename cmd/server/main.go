package main

import "fmt"

type App struct{}

func (app *App) Run() error {
	fmt.Println("Start App")
	return nil
}

func main() {
	fmt.Println("Belajar Go Lagi")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println(err)
	}
}
