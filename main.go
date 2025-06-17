package main

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Data string `json:"data"`
}

func helloHandler(c *gin.Context) {
	response := Response{Data: "Hello World!"}
	c.JSON(200, response)
}

func main() {
	// Инициализируем роутер Gin
	r := gin.Default()

	// Определяем маршрут /hello
	r.GET("/hello", helloHandler)

	// Запускаем сервер на порту 8080
	r.Run(":8080")
}
