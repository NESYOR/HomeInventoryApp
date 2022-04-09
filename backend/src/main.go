package main

import (
	"./api/v1"
	"./constants"
	"fmt"
	"github.com/gofiber/fiber"
	logging "github.com/gofiber/fiber/middleware/logger"
	"../db/infrastructure"
	"log"
)

func main() {
	app:=fiber.New()
	app.Use(logging.New(logging.Config{
		Format: "${time} ${ip} ${pid} ${locals:requestid} ${status} - ${method} ${path}\n"}))
	api:=app.Group("/api")
	version1:=api.Group("/v1")
	version1.Get("/",v1.SayHello)
	version1.Get("/users",v1.GetallUsers)
	version1.Get("/user/:email",v1.GetUser)
	version1.Post("/user",v1.PostUser)
	app.Listen(":3030")
	log.Println("Trying to connect to database")
	go infrastructure.InitConnection()
	fmt.Println(constants.TableNames)
	fmt.Println(constants.Messages["Error"])
}
