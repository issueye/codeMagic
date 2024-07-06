package service

import (
	"github.com/issueye/code_magic/backend/common/model"
	"github.com/issueye/code_magic/backend/common/service"
	"github.com/issueye/code_magic/backend/repository"
	"gorm.io/gorm"
)

type Template struct {
	service.BaseService
}

func (owner *Template) SetBase(base service.BaseService) {
	owner.BaseService = base
}

// 创建数据
func (owner *Template) Create(data *repository.RequestCreateTemplate) error {
	info := model.NewCodeTemplate(&data.CodeTemplateBase)
	return owner.GetDB().Model(&model.CodeTemplate{}).Create(info).Error
}

// 通过ID 获取数据
func (owner *Template) Get(id string) (*model.CodeTemplate, error) {
	info := &model.CodeTemplate{}
	return info, owner.GetDB().Model(&model.CodeTemplate{}).Where("ID = ?", id).First(info).Error
}

// 创建数据
func (owner *Template) Modify(data *repository.RequestModifyTemplate) error {
	updateMap := make(map[string]any)
	updateMap["title"] = data.Title
	updateMap["file_name"] = data.FileName
	updateMap["file_type"] = data.FileType
	updateMap["mark"] = data.Mark
	return owner.GetDB().Model(&model.CodeTemplate{}).Where("ID = ?", data.ID).Updates(&updateMap).Error
}

// 创建数据
func (owner *Template) Delete(id string) error {
	return owner.GetDB().Delete(&model.CodeTemplate{}, id).Error
}

// 获取数据列表
func (owner *Template) List(data *model.Page[repository.RequestTemplateQuery]) ([]*model.CodeTemplate, error) {
	list := make([]*model.CodeTemplate, 0)
	err := service.Filter(owner, model.CodeTemplate{}.TableName(), data, &list, func(db *gorm.DB) (*gorm.DB, error) {

		if data.Condition.Condition != "" {
			db = db.Where("(title like ? file_name like ? file_type like ? or mark like ?)", "%"+data.Condition.Condition+"%", "%"+data.Condition.Condition+"%")
		}

		return db, nil
	})

	return list, err
}
