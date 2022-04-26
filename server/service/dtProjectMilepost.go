package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDtProjectMilepost
//@description: 创建DtProjectMilepost记录
//@param: dtProjectMilepost model.DtProjectMilepost
//@return: err error

func CreateDtProjectMilepost(dtProjectMilepost model.DtProjectMilepost) (err error) {
	err = global.GVA_DB.Create(&dtProjectMilepost).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDtProjectMilepost
//@description: 删除DtProjectMilepost记录
//@param: dtProjectMilepost model.DtProjectMilepost
//@return: err error

func DeleteDtProjectMilepost(dtProjectMilepost model.DtProjectMilepost) (err error) {
	err = global.GVA_DB.Delete(&dtProjectMilepost).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDtProjectMilepostByIds
//@description: 批量删除DtProjectMilepost记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDtProjectMilepostByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.DtProjectMilepost{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDtProjectMilepost
//@description: 更新DtProjectMilepost记录
//@param: dtProjectMilepost *model.DtProjectMilepost
//@return: err error

func UpdateDtProjectMilepost(dtProjectMilepost model.DtProjectMilepost) (err error) {
	err = global.GVA_DB.Save(&dtProjectMilepost).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDtProjectMilepost
//@description: 根据id获取DtProjectMilepost记录
//@param: id uint
//@return: err error, dtProjectMilepost model.DtProjectMilepost

func GetDtProjectMilepost(id uint) (err error, dtProjectMilepost model.DtProjectMilepost) {
	err = global.GVA_DB.Where("id = ?", id).First(&dtProjectMilepost).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDtProjectMilepostInfoList
//@description: 分页获取DtProjectMilepost记录
//@param: info request.DtProjectMilepostSearch
//@return: err error, list interface{}, total int64

func GetDtProjectMilepostInfoList(info request.DtProjectMilepostSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.DtProjectMilepost{})
	var dtProjectMileposts []model.DtProjectMilepost
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&dtProjectMileposts).Error
	return err, dtProjectMileposts, total
}
