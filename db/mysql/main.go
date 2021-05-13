package main

import (
	"log"
	"study/db/mysql/src/boot"
	"study/db/mysql/src/config"
)

func main() {
	boot.GinBoot()
	boot.DBBoot()

	if err := <-config.ErrorChan; err != nil {
		log.Println(err)
	}
}
