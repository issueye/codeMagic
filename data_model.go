package main

import (
	"context"

	"github.com/issueye/code_magic/backend/common/model"
	commonService "github.com/issueye/code_magic/backend/common/service"
	"github.com/issueye/code_magic/backend/repository"
	"github.com/issueye/code_magic/backend/service"
)

// 数据模型
type DataModel struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewDataModel() *DataModel {
	return &DataModel{}
}

// startup is called at application startup
func (a *DataModel) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a *DataModel) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *DataModel) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *DataModel) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// 创建数据模型
func (lc *DataModel) CreateDataModel(data *repository.RequestCreateDataModel) error {
	srv := commonService.NewService(&service.DataModel{})
	return srv.Create(data)
}

func (lc *DataModel) GetDataModelList(condition string, page, size int) ([]*model.DataModel, error) {
	srv := commonService.NewService(&service.DataModel{})
	qry := model.NewPage(repository.RequestDataModelQuery{})
	qry.Condition.Condition = condition
	qry.PageNum = int64(page)
	qry.PageSize = int64(size)

	return srv.List(qry)
}

func (lc *DataModel) GetDataModelInfo(modelId string) ([]*model.ModelInfo, error) {
	srv := commonService.NewService(&service.DataModel{})
	return srv.GetModelInfo(modelId)
}

func (lc *DataModel) SaveModelInfo(modelId string, data []*repository.RequestModelInfoSave) error {
	srv := commonService.NewService(&service.DataModel{})
	return srv.SaveModelInfo(modelId, data)
}

// 创建数据模型
func (lc *DataModel) DeleteDataModel(id string) error {
	srv := commonService.NewService(&service.DataModel{})
	return srv.Delete(id)
}

// 创建数据模型
func (lc *DataModel) ModifyDataModel(data *repository.RequestModifyDataModel) error {
	srv := commonService.NewService(&service.DataModel{})
	return srv.Modify(data)
}
