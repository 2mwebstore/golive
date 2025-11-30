package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Serve HLS files
    r.Static("/streams", "./streams")

    // Serve static page
    r.Static("/static", "./static")
    r.GET("/", func(c *gin.Context) {
        c.File("./static/index.html")
    })

    // Railway assigns dynamic port
    port := "8080"
    if p := getenv("PORT"); p != "" {
        port = p
    }

    r.Run(":" + port)
}

// Simple getenv helper
func getenv(key string) string {
    if v := System.Getenv(key); v != "" {
        return v
    }
    return ""
}
