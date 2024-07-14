package repository

import "github.com/issueye/code_magic/backend/common/model"

type ModifyScheme struct {
	model.Scheme
}

type CreateScheme struct {
	model.Scheme
}

// 查询请求数据
type QryScheme struct {
	Condition string   `json:"condition"` // 条件
	Ids       []string `json:"ids"`       // 模板id
}
