package main

import "fmt"
import "github.com/gin-gonic/gin"

func main() {
	fmt.Println("Hello, world.")
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})

	})
	r.Run()
}
