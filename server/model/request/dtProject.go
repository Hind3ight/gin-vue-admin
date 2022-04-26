package request

import "gin-vue-admin/model"

type DtProjectSearch struct {
	model.DtProject
	PageInfo
}
