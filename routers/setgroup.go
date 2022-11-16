package routers

import "github.com/gin-gonic/gin"

// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/status.php", status)
	r.GET("/remote.php/dav", remote_dav)
	return r
}
