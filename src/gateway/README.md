## 1. Init Project
```bash
go mod init github.com/ducnguyen96/ducnguyen96.xyz
```

## 2. Install Gin
```bash
go get -u github.com/gin-gonic/gin
```

## 3. A server
```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```