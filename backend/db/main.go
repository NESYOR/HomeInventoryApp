package main

import (
	"log"
	"./migrations"
)

func main() {
	log.Println("initializing DB")
	migrations.Initdb()

}
