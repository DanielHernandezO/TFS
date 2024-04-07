package handler

import "github.com/gin-gonic/gin"

type PingHandller interface {
	Ping(c *gin.Context)
}

type pingHandler struct{}

func NewPingHanlder() *pingHandler {
	return &pingHandler{}
}

func (p *pingHandler) Ping(c *gin.Context) {
	c.IndentedJSON(200, "pong")
}
