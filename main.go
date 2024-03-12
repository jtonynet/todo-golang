package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TodoDatabase interface {
	RetrieveTodoList() []Todo
}

type InMemoryTodoDatabase struct {
}

func NewImMemoryTodoDatabase() InMemoryTodoDatabase {
	return InMemoryTodoDatabase{}
}

func (imdb *InMemoryTodoDatabase) RetrieveTodoList() []Todo {
	var todos []Todo

	todos = append(todos, Todo{UUID: uuid.New(), Title: "Title 1"})
	todos = append(todos, Todo{UUID: uuid.New(), Title: "Title 2"})
	todos = append(todos, Todo{UUID: uuid.New(), Title: "Title 3"})

	return todos
}

type Todo struct {
	UUID  uuid.UUID `json:"id"`
	Title string    `json:"tile"`
}

type TodosResponse struct {
	Items []Todo `json:"items"`
}

func main() {
	r := gin.Default()

	r.GET("/todos", func(c *gin.Context) {
		imMemoryDB := NewImMemoryTodoDatabase()

		response := TodosResponse{
			Items: imMemoryDB.RetrieveTodoList(),
		}

		c.JSON(http.StatusOK, response)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
