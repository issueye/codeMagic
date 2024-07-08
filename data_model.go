package main

import (
	"github.com/issueye/code_magic/backend/common/model"
	commonService "github.com/issueye/code_magic/backend/common/service"
	"github.com/issueye/code_magic/backend/logic"
	"github.com/issueye/code_magic/backend/repository"
	"github.com/issueye/code_magic/backend/service"
)

// 数据模型
type DataModel struct {
	BaseApp
}

// NewApp creates a new App application struct
func NewDataModel() *DataModel {
	return &DataModel{}
}

// 创建数据模型
func (lc *DataModel) Create(data *repository.RequestCreateDataModel) error {
	srv := commonService.NewService(&service.DataModel{})
	return srv.Create(data)
}

func (lc *DataModel) List(condition string, page, size int) ([]*model.DataModel, error) {
	srv := commonService.NewService(&service.DataModel{})
	qry := model.NewPage(repository.RequestDataModelQuery{})
	qry.Condition.Condition = condition
	qry.PageNum = int64(page)
	qry.PageSize = int64(size)

	return srv.List(qry)
}

func (lc *DataModel) GetModelInfo(modelId string) ([]*model.ModelInfo, error) {
	srv := commonService.NewService(&service.DataModel{})
	return srv.GetModelInfo(modelId)
}

func (lc *DataModel) SaveModelInfo(modelId string, data []*repository.RequestModelInfoSave) error {
	srv := commonService.NewService(&service.DataModel{})
	return srv.SaveModelInfo(modelId, data)
}

// 创建数据模型
func (lc *DataModel) Delete(id string) error {
	srv := commonService.NewService(&service.DataModel{})
	return srv.Delete(id)
}

// 创建数据模型
func (lc *DataModel) Modify(data *repository.RequestModifyDataModel) error {
	srv := commonService.NewService(&service.DataModel{})
	return srv.Modify(data)
}

// 运行代码
func (lc *DataModel) RunCode(dmId string) error {
	err := logic.RunCode(dmId)
	if err != nil {
		return err
	}

	return nil
}
