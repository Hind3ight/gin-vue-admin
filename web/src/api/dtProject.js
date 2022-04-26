import service from '@/utils/request'

// @Tags DtProject
// @Summary 创建DtProject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DtProject true "创建DtProject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dtProject/createDtProject [post]
export const createDtProject = (data) => {
     return service({
         url: "/dtProject/createDtProject",
         method: 'post',
         data
     })
 }


// @Tags DtProject
// @Summary 删除DtProject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DtProject true "删除DtProject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dtProject/deleteDtProject [delete]
 export const deleteDtProject = (data) => {
     return service({
         url: "/dtProject/deleteDtProject",
         method: 'delete',
         data
     })
 }

// @Tags DtProject
// @Summary 删除DtProject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DtProject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dtProject/deleteDtProject [delete]
 export const deleteDtProjectByIds = (data) => {
     return service({
         url: "/dtProject/deleteDtProjectByIds",
         method: 'delete',
         data
     })
 }

// @Tags DtProject
// @Summary 更新DtProject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DtProject true "更新DtProject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dtProject/updateDtProject [put]
 export const updateDtProject = (data) => {
     return service({
         url: "/dtProject/updateDtProject",
         method: 'put',
         data
     })
 }


// @Tags DtProject
// @Summary 用id查询DtProject
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DtProject true "用id查询DtProject"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dtProject/findDtProject [get]
 export const findDtProject = (params) => {
     return service({
         url: "/dtProject/findDtProject",
         method: 'get',
         params
     })
 }


// @Tags DtProject
// @Summary 分页获取DtProject列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取DtProject列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dtProject/getDtProjectList [get]
 export const getDtProjectList = (params) => {
     return service({
         url: "/dtProject/getDtProjectList",
         method: 'get',
         params
     })
 }