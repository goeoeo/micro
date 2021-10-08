package swagger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opencensus.io/plugin/ochttp"
	"golang.org/x/tools/godoc/vfs"
	"golang.org/x/tools/godoc/vfs/httpfs"
	"golang.org/x/tools/godoc/vfs/mapfs"

	staticSpec "github.com/phpdi/micro/api/generate/swagger/spec"
	staticSwaggerUI "github.com/phpdi/micro/api/generate/swagger/swagger-ui"
)

func Run() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	swaggerGroup := r.Group("swagger-ui")

	swaggerGroup.OPTIONS("/*options_support", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNoContent)
		return
	})
	swaggerGroup.GET("/*filepath", gin.WrapH(handleSwagger()))

	address := fmt.Sprintf(":%s", "8001")
	logrus.Infof("swagger run at http://%s", address)
	err := http.ListenAndServe(address, &ochttp.Handler{Handler: r})
	if err != nil {
		panic(err)
	}
}

func handleSwagger() http.Handler {
	ns := vfs.NameSpace{}
	ns.Bind("/", mapfs.New(staticSwaggerUI.Files), "/", vfs.BindReplace)
	ns.Bind("/", mapfs.New(staticSpec.Files), "/", vfs.BindBefore)
	return http.StripPrefix("/swagger-ui", http.FileServer(httpfs.New(ns)))
}
