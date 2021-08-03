package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	BuildVersion string = ""
	BuildTime    string = ""
)

const ContextRoot = "go-rest-api/v1"

func route(router *gin.Engine, route string) *gin.RouterGroup {
	fr := fmt.Sprintf("%s/%v", ContextRoot, route)
	return router.Group(fr)
}

type Route interface {
	Routes(router *gin.Engine) error
}