package main

import (
    "os"
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

    // Get Railway port
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    r.Run(":" + port)
}
