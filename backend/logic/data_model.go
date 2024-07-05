package logic

import (
	commonService "github.com/issueye/code_magic/backend/common/service"
	"github.com/issueye/code_magic/backend/repository"
	"github.com/issueye/code_magic/backend/service"
)

// 数据模型
type DataMode struct {
}

// 创建数据模型
func (lc *DataMode) CreateDataModel(data *repository.RequestCreateDataModel) error {
	srv := commonService.NewService(&service.DataModel{})
	return srv.Create(data)
}
