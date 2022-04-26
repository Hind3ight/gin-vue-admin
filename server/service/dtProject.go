package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDtProject
//@description: 创建DtProject记录
//@param: dtProject model.DtProject
//@return: err error

func CreateDtProject(dtProject model.DtProject) (err error) {
	err = global.GVA_DB.Create(&dtProject).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDtProject
//@description: 删除DtProject记录
//@param: dtProject model.DtProject
//@return: err error

func DeleteDtProject(dtProject model.DtProject) (err error) {
	err = global.GVA_DB.Delete(&dtProject).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteDtProjectByIds
//@description: 批量删除DtProject记录
//@param: ids request.IdsReq
//@return: err error

func DeleteDtProjectByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.DtProject{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDtProject
//@description: 更新DtProject记录
//@param: dtProject *model.DtProject
//@return: err error

func UpdateDtProject(dtProject model.DtProject) (err error) {
	err = global.GVA_DB.Save(&dtProject).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDtProject
//@description: 根据id获取DtProject记录
//@param: id uint
//@return: err error, dtProject model.DtProject

func GetDtProject(id uint) (err error, dtProject model.DtProject) {
	err = global.GVA_DB.Where("id = ?", id).First(&dtProject).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDtProjectInfoList
//@description: 分页获取DtProject记录
//@param: info request.DtProjectSearch
//@return: err error, list interface{}, total int64

func GetDtProjectInfoList(info request.DtProjectSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.DtProject{})
	var dtProjects []model.DtProject
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&dtProjects).Error
	return err, dtProjects, total
}
