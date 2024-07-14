package model

import (
	"strconv"

	"github.com/issueye/code_magic/backend/pkg/utils"
)

type Scheme struct {
	Base
	SchemeBase
}

type SchemeBase struct {
	Code       string `binding:"required" label:"编码" gorm:"column:code;size:100;comment:编码;" json:"code"`                // 编码
	Title      string `binding:"required" label:"名称" gorm:"column:title;size:300;comment:标题;" json:"title"`              // 标题
	Name       string `binding:"required" label:"名称" gorm:"column:name;size:300;comment:名称;" json:"name"`                // 名称
	Level      int    `binding:"required" label:"等级" gorm:"column:level;size:10;comment:等级;" json:"level"`               // 等级
	ParentCode string `binding:"required" label:"父编码" gorm:"column:parent_code;size:100;comment:父编码;" json:"parentCode"` // 父编码
	NodeType   int    `binding:"required" label:"节点类型" gorm:"column:node_type;size:10;comment:节点类型;" json:"nodeType"`    // 节点类型 0 根节点 1 文件夹节点 2 文件节点
}

func (Scheme) TableName() string {
	return "scheme"
}

func (Scheme) Copy(base *SchemeBase) *Scheme {
	scheme := &Scheme{
		Base:       Base{},
		SchemeBase: *base,
	}

	scheme.ID = strconv.FormatInt(utils.GenID(), 10)
	return scheme
}

func NewScheme(base *SchemeBase) *Scheme {
	db := &Scheme{
		SchemeBase: *base,
	}
	db.ID = strconv.FormatInt(utils.GenID(), 10)

	return db
}
