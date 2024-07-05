package model

import (
	"strconv"

	"github.com/issueye/code_magic/backend/pkg/utils"
)

// 数据模型
type DataModel struct {
	Base
	DataModelBase
}

func NewDataModel(dmBase *DataModelBase) *DataModel {
	db := &DataModel{
		DataModelBase: *dmBase,
	}
	db.ID = strconv.FormatInt(utils.GenID(), 10)

	return db
}

type DataModelBase struct {
	Title    string `binding:"required" label:"名称" gorm:"column:title;size:300;comment:标题;" json:"title"`        // 标题
	MakeType int    `binding:"required" label:"类型" gorm:"column:make_type;type:int;comment:类型;" json:"makeType"` // 类型
	Mark     string `binding:"required" label:"备注" gorm:"column:mark;size:300;comment:备注;" json:"mark"`          // 备注
}

// TableName
// 表名称
func (DataModel) TableName() string {
	return "data_model"
}
