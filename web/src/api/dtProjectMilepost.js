import service from '@/utils/request'

// @Tags DtProjectMilepost
// @Summary 创建DtProjectMilepost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DtProjectMilepost true "创建DtProjectMilepost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dtProjectMilepost/createDtProjectMilepost [post]
export const createDtProjectMilepost = (data) => {
     return service({
         url: "/dtProjectMilepost/createDtProjectMilepost",
         method: 'post',
         data
     })
 }


// @Tags DtProjectMilepost
// @Summary 删除DtProjectMilepost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DtProjectMilepost true "删除DtProjectMilepost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dtProjectMilepost/deleteDtProjectMilepost [delete]
 export const deleteDtProjectMilepost = (data) => {
     return service({
         url: "/dtProjectMilepost/deleteDtProjectMilepost",
         method: 'delete',
         data
     })
 }

// @Tags DtProjectMilepost
// @Summary 删除DtProjectMilepost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DtProjectMilepost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dtProjectMilepost/deleteDtProjectMilepost [delete]
 export const deleteDtProjectMilepostByIds = (data) => {
     return service({
         url: "/dtProjectMilepost/deleteDtProjectMilepostByIds",
         method: 'delete',
         data
     })
 }

// @Tags DtProjectMilepost
// @Summary 更新DtProjectMilepost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DtProjectMilepost true "更新DtProjectMilepost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dtProjectMilepost/updateDtProjectMilepost [put]
 export const updateDtProjectMilepost = (data) => {
     return service({
         url: "/dtProjectMilepost/updateDtProjectMilepost",
         method: 'put',
         data
     })
 }


// @Tags DtProjectMilepost
// @Summary 用id查询DtProjectMilepost
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DtProjectMilepost true "用id查询DtProjectMilepost"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dtProjectMilepost/findDtProjectMilepost [get]
 export const findDtProjectMilepost = (params) => {
     return service({
         url: "/dtProjectMilepost/findDtProjectMilepost",
         method: 'get',
         params
     })
 }


// @Tags DtProjectMilepost
// @Summary 分页获取DtProjectMilepost列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取DtProjectMilepost列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dtProjectMilepost/getDtProjectMilepostList [get]
 export const getDtProjectMilepostList = (params) => {
     return service({
         url: "/dtProjectMilepost/getDtProjectMilepostList",
         method: 'get',
         params
     })
 }