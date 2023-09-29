package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()

	f, _ := os.Create(fmt.Sprintf("%s-%s.log", os.Getenv("GOAUTH_LOG_PATH"), strconv.FormatInt(time.Now().Unix(), 10)))
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	oauth2 := r.Group("/auth2")
	oauth1 := r.Group("/auth1")

	oauth2.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version": "oauth2",
		})
	})

	oauth1.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version": "oauth1",
		})
	})

	r.Run("0.0.0.0:8080")
}
