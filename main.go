package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var InMemoryDB TodoDatabase

type TodoDatabase interface {
	RetrieveTodoList() []Todo
	CreateTodo(todo Todo)
}

type InMemoryTodoDatabase struct {
	todoList map[uuid.UUID]Todo
}

func NewImMemoryTodoDatabase() InMemoryTodoDatabase {
	result := InMemoryTodoDatabase{}
	result.todoList = make(map[uuid.UUID]Todo)
	return result
}

func (imdb *InMemoryTodoDatabase) RetrieveTodoList() []Todo {
	todos := make([]Todo, 0)

	for _, todo := range imdb.todoList {
		todos = append(todos, todo)
	}

	return todos
}

func (imdb *InMemoryTodoDatabase) CreateTodo(todo Todo) {
	imdb.todoList[todo.UUID] = todo
}

type Todo struct {
	UUID  uuid.UUID `json:"id"`
	Title string    `json:"tile"`
}

type TodosResponse struct {
	Items []Todo `json:"items"`
}

func main() {
	InMemoryDB := NewImMemoryTodoDatabase()

	r := gin.Default()

	r.GET("/todos", func(c *gin.Context) {
		response := TodosResponse{
			Items: InMemoryDB.RetrieveTodoList(),
		}

		c.JSON(http.StatusOK, response)
	})

	r.POST("/todos", func(c *gin.Context) {

		//FINALIDADE DE EXEMPLO
		uniqueID := uuid.New()
		t := Todo{UUID: uniqueID, Title: "Title 1"}

		InMemoryDB.CreateTodo(t)
		c.JSON(http.StatusOK, gin.H{"msg": "create"})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
