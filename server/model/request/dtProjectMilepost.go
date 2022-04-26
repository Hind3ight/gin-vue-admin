package request

import "gin-vue-admin/model"

type DtProjectMilepostSearch struct {
	model.DtProjectMilepost
	PageInfo
}
