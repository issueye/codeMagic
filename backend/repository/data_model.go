package repository

import "github.com/issueye/code_magic/backend/common/model"

type RequestCreateDataModel struct {
	model.DataModelBase
}

type RequestModifyDataModel struct {
	model.DataModel
}

// 查询请求数据
type RequestDataModelList struct {
	Condition string `json:"condition"` // 条件
}
