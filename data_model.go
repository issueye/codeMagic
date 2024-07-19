package main

import (
	"context"

	"github.com/issueye/code_magic/backend/common/model"
	commonService "github.com/issueye/code_magic/backend/common/service"
	"github.com/issueye/code_magic/backend/global"
	"github.com/issueye/code_magic/backend/logic"
	"github.com/issueye/code_magic/backend/repository"
	"github.com/issueye/code_magic/backend/service"
)

var dataModel *DataModel

// 数据模型
type DataModel struct {
	Ctx context.Context
}

// NewApp creates a new App application struct
func GetDataModel() *DataModel {
	if dataModel == nil {
		dataModel = &DataModel{Ctx: context.Background()}
	}

	return dataModel
}

// 创建数据模型
func (lc *DataModel) Create(data *repository.CreateDataModel) error {
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
	global.Log.Info("开启事务")
	srv.OpenTx()
	var err error
	defer func() {
		if err != nil {
			global.Log.Errorf("保存模型信息失败 数据回滚: %s", err)
			srv.Rollback()
			return
		}
		global.Log.Info("保存模型信息成功 数据提交")
		srv.Commit()
	}()

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

// 获取变量列表
func (lc *DataModel) GetVariableList(modelId string) ([]*model.Variable, error) {
	return new(logic.DataModel).GetVariableList(modelId)
}

// 删除变量
func (lc *DataModel) DeleteVariable(modelId string, variableId string) error {
	return new(logic.DataModel).DeleteVariable(modelId, variableId)
}

// 保存变量
func (lc *DataModel) SaveVariable(modelId string, variables []*model.VariableBase) error {
	return new(logic.DataModel).AddOrUpdateVariable(modelId, variables)
}

// 运行代码
func (lc *DataModel) RunCode(dmId string) error {
	err := logic.NewCodeLogic(lc.Ctx).RunCode(dmId, false, "")
	if err != nil {
		return err
	}

	return nil
}

func (lc *DataModel) TestRunCode(dmId string, tpId string) error {
	err := logic.NewCodeLogic(lc.Ctx).RunCode(dmId, true, tpId)
	if err != nil {
		return err
	}

	return nil
}
