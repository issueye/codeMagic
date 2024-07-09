package repository

import "github.com/issueye/code_magic/backend/common/model"

type RequestModifyTemplate struct {
	model.CodeTemplate
}

type RequestCreateTemplate struct {
	model.CodeTemplateBase
}

// 查询请求数据
type RequestTemplateQuery struct {
	Condition string   `json:"condition"` // 条件
	Ids       []string `json:"ids"`       // 模板id
}
