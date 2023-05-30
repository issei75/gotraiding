package main

import (
	"fmt"
	"log"

	"example.com/gotrading/config"
)

func main() {
	log.Println("test")
	fmt.Println(config.Config.ApiKey)
}
