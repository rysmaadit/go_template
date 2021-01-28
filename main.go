package main

import (
	"example.com/m/app"
	"example.com/m/cmd"
	"os"
)

func main() {
	c := cmd.NewCli(os.Args)
	app.NewApp()
	c.Run()
}
