package main

import (
	"fmt"
	"hello/config"
	"log"
)

func main() {
	// configリストを読み込んで表示する
	// configパッケージのConfigという変数のPort
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	log.Panicln("test")
}
