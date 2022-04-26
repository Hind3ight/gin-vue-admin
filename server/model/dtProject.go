// 自动生成模板DtProject
package model

import (
	"gin-vue-admin/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type DtProject struct {
	global.GVA_MODEL
	Name          string     `json:"name" form:"name" gorm:"column:name;comment:项目名称;type:longtext"`
	Profile       string     `json:"profile" form:"profile" gorm:"column:profile;comment:项目简介;type:longtext"`
	ProjectNumber string     `json:"projectNumber" form:"projectNumber" gorm:"column:project_number;comment:合同备案编号;type:longtext"`
	StartDate     *time.Time `json:"startDate" form:"startDate" gorm:"column:start_date;comment:开工日期;type:date;"`
	EndDate       *time.Time `json:"endDate" form:"endDate" gorm:"column:end_date;comment:完工日期;type:date;"`
	JsId          int        `json:"jsId" form:"jsId" gorm:"column:js_id;comment:建设单位"`
	KcId          int        `json:"kcId" form:"kcId" gorm:"column:kc_id;comment:勘察单位"`
	SjId          int        `json:"sjId" form:"sjId" gorm:"column:sj_id;comment:设计单位"`
	SgId          int        `json:"sgId" form:"sgId" gorm:"column:sg_id;comment:施工单位"`
	JlId          int        `json:"jlId" form:"jlId" gorm:"column:Jl_id;comment:监理单位"`
	Source        string     `json:"source" form:"source" gorm:"column:source;comment:数据来源;type:longtext"`
}

func (DtProject) TableName() string {
	return "dt_project"
}
