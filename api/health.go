package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Health() Route {
	return health{url: "health"}
}

type health struct {
	url string
}

func (h health) Routes(router *gin.Engine) error {
	handler, err := newHealthHandler()
	if err!=nil{
		return err
	}
	r := route(router,h.url)

	r.GET("/ping", handler.health)
	r.GET("/version", handler.version)

	return nil
}



type healthHandler struct {

}

func newHealthHandler() (*healthHandler,error)  {
	return &healthHandler{},nil
}

func (h healthHandler) health(c *gin.Context)  {
	c.JSON(http.StatusOK, "Ping Successfully")
}

func (h healthHandler) version(c *gin.Context) {
	c.JSON(http.StatusOK, BuildVersion)
}