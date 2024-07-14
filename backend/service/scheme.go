package service

import (
	"github.com/issueye/code_magic/backend/common/model"
	"github.com/issueye/code_magic/backend/common/service"
	"github.com/issueye/code_magic/backend/repository"
	"gorm.io/gorm"
)

type Scheme struct {
	service.BaseService
}

func (owner *Scheme) SetBase(base service.BaseService) {
	owner.BaseService = base
}

// 创建数据
func (owner *Scheme) Create(data *repository.CreateScheme) error {
	info := model.NewScheme(&data.SchemeBase)
	return owner.GetDB().Model(&model.Scheme{}).Create(info).Error
}

// 通过ID 获取数据
func (owner *Scheme) Get(id string) (*model.Scheme, error) {
	info := &model.Scheme{}
	return info, owner.GetDB().Model(&model.Scheme{}).Where("ID = ?", id).First(info).Error
}

// 通过ID 获取数据
func (owner *Scheme) Gets(ids []string) ([]*model.Scheme, error) {
	info := make([]*model.Scheme, 0)
	err := owner.GetDB().Model(&model.Scheme{}).Where("ID in (?)", ids).Find(&info).Error
	return info, err
}

// 创建数据
func (owner *Scheme) Modify(data *repository.ModifyScheme) error {
	updateMap := make(map[string]any)
	updateMap["title"] = data.Title
	return owner.GetDB().Model(&model.Scheme{}).Where("ID = ?", data.ID).Updates(&updateMap).Error
}

// 创建数据
func (owner *Scheme) Delete(id string) error {
	return owner.GetDB().Delete(&model.Scheme{}, id).Error
}

// 获取数据列表
func (owner *Scheme) List(data *model.Page[repository.QryScheme]) ([]*model.Scheme, error) {
	list := make([]*model.Scheme, 0)
	err := service.Filter(owner, model.Scheme{}.TableName(), data, &list, func(db *gorm.DB) (*gorm.DB, error) {

		if data.Condition.Condition != "" {
			db = db.Where("(title like ? file_name like ? file_type like ? or mark like ?)", "%"+data.Condition.Condition+"%", "%"+data.Condition.Condition+"%")
		}

		if len(data.Condition.Ids) > 0 {
			db = db.Where("id in (?)", data.Condition.Ids)
		}

		return db, nil
	})

	return list, err
}
