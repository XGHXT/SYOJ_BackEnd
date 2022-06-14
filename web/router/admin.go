package router

import (
	"github.com/gin-gonic/gin"
	"SYOJ_BackEnd/web/api/super_admin"
)

// InitAdminRouter 初始化具体的路由组
func InitAdminRouter(router *gin.RouterGroup) {
	uGroup := router.Group("/admin")

	{
		uGroup.POST("/superadmin", super_admin.Create)
	}
}
