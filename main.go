package main

import (
	"fmt"
	"hello/app/controllers"
	"hello/app/models"
)

func main() {
	// init関数を呼びだす
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
