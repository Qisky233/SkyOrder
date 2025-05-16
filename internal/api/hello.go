package api

import "github.com/gin-gonic/gin"

type HelloHandler struct{}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (h *HelloHandler) Hello(c *gin.Context) {
	c.String(200, "Hello, World!")
}
