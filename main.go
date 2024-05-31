package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Создаем роутер Gin
	router := gin.Default()
	port := os.Getenv("PORT")
	name := os.Getenv("TEST_NAME")

	// Регистрируем обработчик для корневого пути
	router.GET("/", func(c *gin.Context) {
		c.String(200, fmt.Sprintf("Hello, World, %s", name))
	})

	// Запускаем сервер на адресе 0.0.0.0:8080
	router.Run(fmt.Sprintf("0.0.0.0:%s", port))
}
