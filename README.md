# RedisRESTAPI

A simple REST API for managing todo items using Go, Fiber, and Redis. This project demonstrates how to create, retrieve, update, and delete todo items stored in Redis through HTTP endpoints.

## Features

- **Create Todo**: Add a new todo item to the Redis store.
- **Get All Todos**: Retrieve a list of all todo items stored in Redis.
- **Get Todo by ID**: Fetch a specific todo item using its ID.
- **Update Todo**: Modify an existing todo item (title and completion status).
- **Delete Todo**: Remove a todo item from Redis by ID.

## Technologies Used

- **Go** (v1.23.3)
- **Fiber** (v2.52.5) - A fast and lightweight web framework for Go.
- **Redis** (v9.7.0) - A high-performance in-memory database used for storing todo items.
- **UUID** - For generating unique identifiers.
- **Gofiber** - Web framework used to handle HTTP requests and routing.

## Getting Started

### Prerequisites

- Install Go (version 1.23.3 or later).
- Install Redis and ensure it's running locally on port `6379`.

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/RedisRESTAPI.git
    cd RedisRESTAPI
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Run the application:

    ```bash
    go run main.go
    ```

   The server will be running on `http://localhost:3000`.

## API Endpoints

### `POST /todos`
- **Description**: Creates a new Todo item.
- **Request Body**:
    ```json
    {
        "title": "Your Todo Title"
    }
    ```
- **Response**:
    ```json
    {
        "id": 1,
        "title": "Your Todo Title",
        "completed": false
    }
    ```

### `GET /todos`
- **Description**: Retrieves a list of all Todo items.
- **Response**:
    ```json
    [
        {
            "id": 1,
            "title": "Your Todo Title",
            "completed": false
        }
    ]
    ```

### `GET /todos/:id`
- **Description**: Retrieves a specific Todo item by ID.
- **Response**:
    ```json
    {
        "id": 1,
        "title": "Your Todo Title",
        "completed": false
    }
    ```

### `PUT /todos/:id`
- **Description**: Updates an existing Todo item by ID.
- **Request Body**:
    ```json
    {
        "title": "Updated Todo Title",
        "completed": true
    }
    ```
- **Response**:
    ```json
    {
        "id": 1,
        "title": "Updated Todo Title",
        "completed": true
    }
    ```

### `DELETE /todos/:id`
- **Description**: Deletes a Todo item by ID.
- **Response**: 200 OK

## Contribution

Feel free to fork this repository and contribute to the development of RedisRESTAPI. Any improvements, bug fixes, and feature additions are welcome.

## License

This project is open-source and available under the [MIT License](LICENSE).
