package main

import (
	"context"

	"github.com/issueye/code_magic/backend/common/model"
	commonService "github.com/issueye/code_magic/backend/common/service"
	"github.com/issueye/code_magic/backend/repository"
	"github.com/issueye/code_magic/backend/service"
)

var dataSource *DataSource

// 数据模型
type DataSource struct {
	Ctx context.Context
}

// NewApp creates a new App application struct
func GeDataSource() *DataSource {
	if dataSource == nil {
		dataSource = &DataSource{Ctx: context.Background()}
	}

	return dataSource
}

// 创建数据模型
func (lc *DataSource) Create(data *repository.CreateDataSource) error {
	srv := commonService.NewService(&service.DataSource{})
	return srv.Create(data)
}

// 创建数据模型
func (lc *DataSource) Delete(id string) error {
	srv := commonService.NewService(&service.DataSource{})
	return srv.Delete(id)
}

// 创建数据模型
func (lc *DataSource) Modify(data *repository.UpdateDataSource) error {
	srv := commonService.NewService(&service.DataSource{})
	return srv.Modify(data)
}

func (lc *DataSource) List(condition string, page, size int) ([]*model.DataSource, error) {
	srv := commonService.NewService(&service.DataSource{})
	qry := model.NewPage(repository.QueryDataSource{})
	qry.Condition.Condition = condition
	qry.PageNum = int64(page)
	qry.PageSize = int64(size)

	return srv.List(qry)
}
