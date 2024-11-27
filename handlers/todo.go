package handlers

import (
	"RedisRESTAPI/redis"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

// CreateTodo creates a new todo item and stores it in Redis
func CreateTodo(c *fiber.Ctx) error {
	// Define request structure for incoming JSON data
	type Request struct {
		Title string `json:"title"` // The title of the todo item
	}
	var req Request
	// Parse the incoming request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Input"}) // If parsing fails, return an error
	}

	// Retrieve the current todo_id from Redis to increment and generate a new unique ID
	todoID := redis.Client.Get(redis.Ctx, "todo_id").Val()
	if todoID == "" {
		// If the todo_id doesn't exist, initialize it with 0
		redis.Client.Set(redis.Ctx, "todo_id", 0, 0)
	}

	// Increment the todo_id in Redis to generate a new ID
	id, err := redis.Client.Incr(redis.Ctx, "todo_id").Result()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate id"}) // If increment fails, return an error
	}

	// Create a new key for the todo item using the generated ID
	key := fmt.Sprintf("todo:%d", id)
	// Store the todo item in Redis using HSet (hash set)
	err = redis.Client.HSet(redis.Ctx, key, map[string]interface{}{
		"title":     req.Title,
		"completed": "0", // Default to 'not completed'
	}).Err()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()}) // If storing in Redis fails, return an error
	}

	// Return the newly created todo item in the response
	return c.Status(201).JSON(fiber.Map{
		"id":        id,
		"title":     req.Title,
		"completed": false,
	})
}

// GetTodos retrieves all todos from Redis
func GetTodos(c *fiber.Ctx) error {
	// Fetch all keys for todo items from Redis
	keys, err := redis.Client.Keys(redis.Ctx, "todo:*").Result()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch keys"}) // If fetching keys fails, return an error
	}

	var todos []map[string]interface{}
	// Loop through all keys to fetch each todo item
	for _, key := range keys {
		todo, err := redis.Client.HGetAll(redis.Ctx, key).Result()
		if err != nil || len(todo) == 0 {
			continue // Skip if fetching the todo fails or if the todo is empty
		}

		// Convert "completed" field to a boolean
		completed := todo["completed"] == "1"

		// Append the todo item to the todos list
		todos = append(todos, map[string]interface{}{
			"id":        key[len("todo:"):],
			"title":     todo["title"],
			"completed": completed,
		})
	}

	// Return all todos in the response
	return c.JSON(todos)
}

// GetTodo retrieves a single todo by its ID from Redis
func GetTodo(c *fiber.Ctx) error {
	// Get the todo ID from the URL parameters
	id := c.Params("id")
	key := fmt.Sprintf("todo:%s", id)

	// Check if the todo item exists in Redis
	exists, err := redis.Client.Exists(redis.Ctx, key).Result()
	if err != nil || exists == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"}) // If not found, return a 404 error
	}

	// Fetch the todo item from Redis
	todo, err := redis.Client.HGetAll(redis.Ctx, key).Result()
	if err != nil || len(todo) == 0 {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch todo"}) // If fetching fails, return an error
	}

	// Convert "completed" field to a boolean
	completed := todo["completed"] == "1"

	// Return the fetched todo item
	return c.JSON(fiber.Map{
		"id":        id,
		"title":     todo["title"],
		"completed": completed,
	})
}

// UpdateTodo updates an existing todo item by its ID
func UpdateTodo(c *fiber.Ctx) error {
	// Get the todo ID from the URL parameters
	id := c.Params("id")
	key := fmt.Sprintf("todo:%s", id)

	// Check if the todo item exists in Redis
	existing, err := redis.Client.Exists(redis.Ctx, key).Result()
	if err != nil || existing == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"}) // If not found, return a 404 error
	}

	// Define request structure for the incoming update data
	type Request struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	var req Request
	// Parse the incoming request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Input"}) // If parsing fails, return an error
	}

	// Convert "completed" field to string (Redis stores as "0" or "1")
	completed := "0"
	if req.Completed {
		completed = "1"
	}

	// Update the todo item in Redis
	err = redis.Client.HSet(redis.Ctx, key, map[string]interface{}{
		"title":     req.Title,
		"completed": completed,
	}).Err()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update todo"}) // If updating fails, return an error
	}

	// Return the updated todo item
	return c.JSON(fiber.Map{"id": id, "title": req.Title, "completed": req.Completed})
}

// DeleteTodo deletes a todo item by its ID
func DeleteTodo(c *fiber.Ctx) error {
	// Get the todo ID from the URL parameters
	id := c.Params("id")
	key := fmt.Sprintf("todo:%s", id)

	// Delete the todo item from Redis
	_, err := redis.Client.Del(redis.Ctx, key).Result()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete todo"}) // If deletion fails, return an error
	}

	// Return success status
	return c.SendStatus(fiber.StatusOK)
}
