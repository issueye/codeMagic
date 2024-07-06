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

type ModelInfoBase struct {
	DataModelId string `binding:"required" label:"模型ID" gorm:"column:data_model_id;size:300;comment:模型ID;" json:"dataModelId"` // 模型ID
	Title       string `binding:"required" label:"名称" gorm:"column:title;size:300;comment:标题;" json:"title"`                   // 标题
	Name        string `binding:"required" label:"标识" gorm:"column:name;size:300;comment:标识;" json:"name"`                     // 标识
	ColumnType  string `binding:"required" label:"类型" gorm:"column:column_type;size:50;comment:类型;" json:"columnType"`         // 类型
	Size        int    `binding:"required" label:"长度" gorm:"column:size;type:int;comment:长度;" json:"size"`                     // 长度
	Mark        string `binding:"required" label:"备注" gorm:"column:mark;size:300;comment:备注;" json:"mark"`                     // 备注
}

type ModelInfo struct {
	Base
	ModelInfoBase
}

func (ModelInfo) TableName() string {
	return "model_info"
}

func NewModelInfo(mib *ModelInfoBase) *ModelInfo {
	db := &ModelInfo{
		ModelInfoBase: *mib,
	}
	db.ID = strconv.FormatInt(utils.GenID(), 10)

	return db
}
