package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDtProjectRouter(Router *gin.RouterGroup) {
	DtProjectRouter := Router.Group("dtProject").Use(middleware.OperationRecord())
	{
		DtProjectRouter.POST("createDtProject", v1.CreateDtProject)             // 新建DtProject
		DtProjectRouter.DELETE("deleteDtProject", v1.DeleteDtProject)           // 删除DtProject
		DtProjectRouter.DELETE("deleteDtProjectByIds", v1.DeleteDtProjectByIds) // 批量删除DtProject
		DtProjectRouter.PUT("updateDtProject", v1.UpdateDtProject)              // 更新DtProject
		DtProjectRouter.GET("findDtProject", v1.FindDtProject)                  // 根据ID获取DtProject
		DtProjectRouter.GET("getDtProjectList", v1.GetDtProjectList)            // 获取DtProject列表
	}
}
