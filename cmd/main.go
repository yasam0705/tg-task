package main

import (
	_ "tg_task/docs"
	"tg_task/pkg/handler"
)

// @title Telegram Bot task
// @version 1.0
// @description Send messages to telegram group or channel

// @host localhost:8080
// @BasePath /

func main() {

	handler.InitRoutes()
}
