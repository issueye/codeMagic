package service

import (
	"github.com/issueye/code_magic/backend/common/model"
	"github.com/issueye/code_magic/backend/common/service"
	"github.com/issueye/code_magic/backend/repository"
	"gorm.io/gorm"
)

type DataModel struct {
	service.BaseService
}

func (owner *DataModel) SetBase(base service.BaseService) {
	owner.BaseService = base
}

// 创建数据
func (owner *DataModel) Create(data *repository.RequestCreateDataModel) error {
	info := model.NewDataModel(&model.DataModelBase{
		Title:    data.Title,
		MakeType: data.MakeType,
		Mark:     data.Mark,
	})

	return owner.GetDB().Model(&model.DataModel{}).Create(info).Error
}

// 创建数据
func (owner *DataModel) Modify(data *repository.RequestModifyDataModel) error {
	updateMap := make(map[string]any)
	updateMap["title"] = data.Title
	updateMap["make_type"] = data.MakeType
	updateMap["mark"] = data.Mark
	return owner.GetDB().Model(&model.DataModel{}).Where("ID = ?", data.ID).Updates(&updateMap).Error
}

// 创建数据
func (owner *DataModel) Delete(id string) error {
	return owner.GetDB().Delete(&model.DataModel{}, id).Error
}

// 获取数据列表
func (owner *DataModel) List(data *model.Page[repository.RequestDataModelList]) ([]*model.DataModel, error) {
	list := make([]*model.DataModel, 0)
	err := service.Filter(owner, model.DataModel{}.TableName(), data, &list, func(db *gorm.DB) (*gorm.DB, error) {

		if data.Condition.Condition != "" {
			db = db.Where("(title like ? or mark like ?)", "%"+data.Condition.Condition+"%", "%"+data.Condition.Condition+"%")
		}

		return db, nil
	})

	return list, err
}
