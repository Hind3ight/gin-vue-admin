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

// @Tags DtProjectMilepost
// @Summary 创建DtProjectMilepost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DtProjectMilepost true "创建DtProjectMilepost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dtProjectMilepost/createDtProjectMilepost [post]
func CreateDtProjectMilepost(c *gin.Context) {
	var dtProjectMilepost model.DtProjectMilepost
	_ = c.ShouldBindJSON(&dtProjectMilepost)
	if err := service.CreateDtProjectMilepost(dtProjectMilepost); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags DtProjectMilepost
// @Summary 删除DtProjectMilepost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DtProjectMilepost true "删除DtProjectMilepost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dtProjectMilepost/deleteDtProjectMilepost [delete]
func DeleteDtProjectMilepost(c *gin.Context) {
	var dtProjectMilepost model.DtProjectMilepost
	_ = c.ShouldBindJSON(&dtProjectMilepost)
	if err := service.DeleteDtProjectMilepost(dtProjectMilepost); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags DtProjectMilepost
// @Summary 批量删除DtProjectMilepost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DtProjectMilepost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dtProjectMilepost/deleteDtProjectMilepostByIds [delete]
func DeleteDtProjectMilepostByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDtProjectMilepostByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags DtProjectMilepost
// @Summary 更新DtProjectMilepost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DtProjectMilepost true "更新DtProjectMilepost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dtProjectMilepost/updateDtProjectMilepost [put]
func UpdateDtProjectMilepost(c *gin.Context) {
	var dtProjectMilepost model.DtProjectMilepost
	_ = c.ShouldBindJSON(&dtProjectMilepost)
	if err := service.UpdateDtProjectMilepost(dtProjectMilepost); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags DtProjectMilepost
// @Summary 用id查询DtProjectMilepost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DtProjectMilepost true "用id查询DtProjectMilepost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dtProjectMilepost/findDtProjectMilepost [get]
func FindDtProjectMilepost(c *gin.Context) {
	var dtProjectMilepost model.DtProjectMilepost
	_ = c.ShouldBindQuery(&dtProjectMilepost)
	if err, redtProjectMilepost := service.GetDtProjectMilepost(dtProjectMilepost.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redtProjectMilepost": redtProjectMilepost}, c)
	}
}

// @Tags DtProjectMilepost
// @Summary 分页获取DtProjectMilepost列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DtProjectMilepostSearch true "分页获取DtProjectMilepost列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dtProjectMilepost/getDtProjectMilepostList [get]
func GetDtProjectMilepostList(c *gin.Context) {
	var pageInfo request.DtProjectMilepostSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDtProjectMilepostInfoList(pageInfo); err != nil {
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
