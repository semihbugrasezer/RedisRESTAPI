package main

import (
	"RedisRESTAPI/handlers"
	"RedisRESTAPI/redis"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	redis.ConnectRedis()

	app := fiber.New()

	//Routes
	app.Post("/todos", handlers.CreateTodo)
	app.Get("/todos", handlers.GetTodos)
	app.Get("/todos/:id", handlers.GetTodo)
	app.Put("/todos/:id", handlers.UpdateTodo)
	app.Delete("/todos/:id", handlers.DeleteTodo)

	log.Fatal(app.Listen(":3000"))

}
