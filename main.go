package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Создаем роутер Gin
	router := gin.Default()

	// Регистрируем обработчик для корневого пути
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// Запускаем сервер на адресе 0.0.0.0:8080
	router.Run("0.0.0.0:8080")
}
