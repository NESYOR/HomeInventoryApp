package main

import (
	"log"

	"./migrations"
)

func main() {
	log.Println("intializing DB")
	migrations.Initdb()

}
