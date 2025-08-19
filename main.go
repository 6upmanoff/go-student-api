package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var students = []Student{
	{ID: 1, Name: "Айдана", Age: 22},
	{ID: 2, Name: "Ержан", Age: 24},
}

func main() {
	r := gin.Default()

	// Получить всех студентов
	r.GET("/students", func(c *gin.Context) {
		c.JSON(http.StatusOK, students)
	})

	// Добавить студента
	r.POST("/students", func(c *gin.Context) {
		var newStudent Student
		if err := c.BindJSON(&newStudent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newStudent.ID = len(students) + 1
		students = append(students, newStudent)
		c.JSON(http.StatusCreated, newStudent)
	})

	// Запустить сервер
	r.Run(":8080")
}
