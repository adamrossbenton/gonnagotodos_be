package actions

import (
    "net/http"
    "github.com/gobuffalo/buffalo"
	"todosbe/models"
)

// TodoIndex default implementation.
func TodoIndex(c buffalo.Context) error {
	// Create array to receive todos
	todos := []models.Todo{}
	// Get all todos from db
	err := models.DB.All(&todos)
	// Handle error
	if err != nil {
		return c.Render(http.StatusOK, r.JSON(err))
	}
	// Return list of todos as json
	return c.Render(http.StatusOK, r.JSON(todos))
}

// TodoShow default implementation.
func TodoShow(c buffalo.Context) error {
	// Get id url param defined in app.go
	id := c.Param("id")
	// Create variable to receive todo
	todo := models.Todo{}
	// Grab todo from db
	err := models.DB.Find(&todo, id)
	// Handle error
	if err != nil {
		return c.Render(http.StatusOK, r.JSON(err))
	}
	// Return todo as json
	return c.Render(http.StatusOK, r.JSON(&todo))
}

// TodoAdd default implementation.
func TodoAdd(c buffalo.Context) error {
	// Get item from url query
	item := c.Param("item")
	// Create new instance of todo
	todo := models.Todo{Item: item}
	// Create todo without running validations
	err := models.DB.Create(&todo)
	// Handle error
	if err != nil {
		return c.Render(http.StatusOK, r.JSON(err))
	}
	// Return new todo as json
	return c.Render(http.StatusOK, r.JSON(todo))
}