package model

import (
	"strconv"

	"github.com/issueye/code_magic/backend/pkg/utils"
)

type Variable struct {
	Base
	VariableBase
}

type VariableBase struct {
	MDId  string `binding:"required" label:"模型ID" gorm:"column:md_id;size:20;comment:模型ID;" json:"md_id"` // 模型ID
	Key   string `binding:"required" label:"键" gorm:"column:key;size:100;comment:键;" json:"key"`          // 键
	Value string `binding:"required" label:"参数值" gorm:"column:value;size:2000;comment:参数值;" json:"value"` // 参数值
	Mark  string `binding:"required" label:"备注" gorm:"column:mark;size:2000;comment:备注;" json:"mark"`     // 备注
}

func (Variable) TableName() string {
	return "variable"
}

func (Variable) Copy(base *VariableBase) *Variable {
	Variable := &Variable{
		Base:         Base{},
		VariableBase: *base,
	}

	Variable.ID = strconv.FormatInt(utils.GenID(), 10)
	return Variable
}

func NewVariable(base *VariableBase) *Variable {
	db := &Variable{
		VariableBase: *base,
	}
	db.ID = strconv.FormatInt(utils.GenID(), 10)

	return db
}
