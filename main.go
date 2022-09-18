package main

import (
	"fmt"
	"hello/app/controllers"
	"hello/app/models"
)

func main() {
	// init関数を呼びだす
	fmt.Println(models.Db)

	// 関数呼び出し
	controllers.StartMainServer()

	// user, _ := models.GetUserByEmail("kohei@kohei.com")
	// fmt.Println(user)

	// session, err := user.CreateSession()
	// if err != nil {
	// 	log.Panicln(err)
	// }
	// fmt.Println(session)

	// valid, _ := session.CheckSession()
	// fmt.Println(valid)
}
