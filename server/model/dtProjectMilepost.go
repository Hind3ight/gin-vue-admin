// 自动生成模板DtProjectMilepost
package model

import (
	"gin-vue-admin/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type DtProjectMilepost struct {
	global.GVA_MODEL
	ProjectId      int        `json:"projectId" form:"projectId" gorm:"column:project_id;comment:项目;type:int;size:10;"`
	Sequence       int        `json:"sequence" form:"sequence" gorm:"column:sequence;comment:排序;type:int;size:10;"`
	Name           string     `json:"name" form:"name" gorm:"column:name;comment:里程碑名称;type:longtext;"`
	PlanStart      *time.Time `json:"planStart" form:"planStart" gorm:"column:plan_start;comment:启动时间;type:date;"`
	PlanEnd        *time.Time `json:"planEnd" form:"planEnd" gorm:"column:plan_end;comment:完成时间;type:date;"`
	Checktime      *time.Time `json:"checktime" form:"checktime" gorm:"column:checktime;comment:验收时间;type:date;"`
	State          string     `json:"state" form:"state" gorm:"column:state;comment:状态;type:longtext;"`
	Remainday      int        `json:"remainday" form:"remainday" gorm:"column:remainday;comment:剩余天数;type:int;size:10;"`
	PracticalStart *time.Time `json:"practicalStart" form:"practicalStart" gorm:"column:practical_start;comment:实际启动时间;type:date;"`
	PracticalEnd   *time.Time `json:"practicalEnd" form:"practicalEnd" gorm:"column:practical_end;comment:实际完成时间;type:date;"`
}

func (DtProjectMilepost) TableName() string {
	return "dt_project_milepost"
}
