package main

import (
	"cach-plan.go/app"
	"cach-plan.go/config"
)

func main()  {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":8080")
}
