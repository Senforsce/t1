package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/senforsce/t1/examples/integration-gin/gintemplrenderer"
)

func main() {
	engine := gin.Default()
	engine.HTMLRender = gintemplrenderer.Default

	// Disable trusted proxy warning.
	engine.SetTrustedProxies(nil)

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", Home())
	})

	engine.GET("/with-ctx", func(c *gin.Context) {
		r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, Home())
		c.Render(http.StatusOK, r)
	})

	engine.Run(":8080")
}
