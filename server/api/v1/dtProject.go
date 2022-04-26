package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags DtProject
// @Summary 创建DtProject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DtProject true "创建DtProject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dtProject/createDtProject [post]
func CreateDtProject(c *gin.Context) {
	var dtProject model.DtProject
	_ = c.ShouldBindJSON(&dtProject)
	if err := service.CreateDtProject(dtProject); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags DtProject
// @Summary 删除DtProject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DtProject true "删除DtProject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dtProject/deleteDtProject [delete]
func DeleteDtProject(c *gin.Context) {
	var dtProject model.DtProject
	_ = c.ShouldBindJSON(&dtProject)
	if err := service.DeleteDtProject(dtProject); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags DtProject
// @Summary 批量删除DtProject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DtProject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dtProject/deleteDtProjectByIds [delete]
func DeleteDtProjectByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDtProjectByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags DtProject
// @Summary 更新DtProject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DtProject true "更新DtProject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dtProject/updateDtProject [put]
func UpdateDtProject(c *gin.Context) {
	var dtProject model.DtProject
	_ = c.ShouldBindJSON(&dtProject)
	if err := service.UpdateDtProject(dtProject); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags DtProject
// @Summary 用id查询DtProject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DtProject true "用id查询DtProject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dtProject/findDtProject [get]
func FindDtProject(c *gin.Context) {
	var dtProject model.DtProject
	_ = c.ShouldBindQuery(&dtProject)
	if err, redtProject := service.GetDtProject(dtProject.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redtProject": redtProject}, c)
	}
}

// @Tags DtProject
// @Summary 分页获取DtProject列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DtProjectSearch true "分页获取DtProject列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dtProject/getDtProjectList [get]
func GetDtProjectList(c *gin.Context) {
	var pageInfo request.DtProjectSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDtProjectInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
