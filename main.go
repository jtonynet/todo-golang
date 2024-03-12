package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
		todo01 := Todo{UUID: uuid.New(), Title: "Title 1"}
		todo02 := Todo{UUID: uuid.New(), Title: "Title 2"}
		todo03 := Todo{UUID: uuid.New(), Title: "Title 3"}

		response := TodosResponse{
			Items: []Todo{todo01, todo02, todo03},
		}

		c.JSON(http.StatusOK, response)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
