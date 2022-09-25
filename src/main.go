package main

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
)

func main() {
  gin.SetMode(gin.ReleaseMode)
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.NoRoute(func(c *gin.Context) {
      c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
  })
  fmt.Println("Starting the application...")
  r.Run("0.0.0.0:8080")
}