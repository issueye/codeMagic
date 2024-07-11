package model

import (
	"strconv"

	"github.com/issueye/code_magic/backend/pkg/utils"
)

// 数据模型
type CodeTemplate struct {
	Base
	CodeTemplateBase
}

// TableName
// 表名称
func (CodeTemplate) TableName() string {
	return "code_template"
}

type CodeTemplateBase struct {
	Title    string `binding:"required" label:"名称" gorm:"column:title;size:300;comment:标题;" json:"title"`            // 标题
	FileName string `binding:"required" label:"文件名" gorm:"column:file_name;size:300;comment:文件名;" json:"fileName"`   // 文件名
	FileType int    `binding:"required" label:"文件类型" gorm:"column:file_type;type:int;comment:文件类型;" json:"fileType"` // 文件类型
	Mark     string `binding:"required" label:"备注" gorm:"column:mark;size:300;comment:备注;" json:"mark"`              // 备注
}

var fileTypeMap = map[int]string{
	1: "go",
	2: "javascript",
	3: "vue",
}

// 获取文件类型
func GetFileType(code int) string {
	return fileTypeMap[code]
}

func NewCodeTemplate(base *CodeTemplateBase) *CodeTemplate {
	db := &CodeTemplate{
		CodeTemplateBase: *base,
	}
	db.ID = strconv.FormatInt(utils.GenID(), 10)

	return db
}
