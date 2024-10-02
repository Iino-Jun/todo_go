package main

import (
	"fmt"
	"go_todo_app/apps/controllers"
	"go_todo_app/apps/models"
)

func main() {
	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/
	fmt.Println(models.Db)

	if err := controllers.StartMainServer(); err != nil {
		fmt.Println("Server failed to start:", err)
	}

}
