// export GIN_MODE=release, 静默输出
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/pingTime", func(c *gin.Context) {
		// JSON serializer is available on gin context
		c.JSON(http.StatusOK, gin.H{
			"serverTime": time.Now().UTC(),
		})
	})
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
