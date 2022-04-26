package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDtProjectMilepostRouter(Router *gin.RouterGroup) {
	DtProjectMilepostRouter := Router.Group("dtProjectMilepost").Use(middleware.OperationRecord())
	{
		DtProjectMilepostRouter.POST("createDtProjectMilepost", v1.CreateDtProjectMilepost)             // 新建DtProjectMilepost
		DtProjectMilepostRouter.DELETE("deleteDtProjectMilepost", v1.DeleteDtProjectMilepost)           // 删除DtProjectMilepost
		DtProjectMilepostRouter.DELETE("deleteDtProjectMilepostByIds", v1.DeleteDtProjectMilepostByIds) // 批量删除DtProjectMilepost
		DtProjectMilepostRouter.PUT("updateDtProjectMilepost", v1.UpdateDtProjectMilepost)              // 更新DtProjectMilepost
		DtProjectMilepostRouter.GET("findDtProjectMilepost", v1.FindDtProjectMilepost)                  // 根据ID获取DtProjectMilepost
		DtProjectMilepostRouter.GET("getDtProjectMilepostList", v1.GetDtProjectMilepostList)            // 获取DtProjectMilepost列表
	}
}
