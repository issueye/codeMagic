package logic

import (
	"github.com/issueye/code_magic/backend/common/model"
	commonService "github.com/issueye/code_magic/backend/common/service"
	"github.com/issueye/code_magic/backend/global"
	"github.com/issueye/code_magic/backend/repository"
	"github.com/issueye/code_magic/backend/service"
)

// 数据模型
type DataModel struct {
}

// 创建数据模型
func (lc *DataModel) CreateDataModel(data *repository.CreateDataModel) error {
	srv := commonService.NewService(&service.DataModel{})
	return srv.Create(data)
}

// 新增或者修改变量
func (lc *DataModel) AddOrUpdateVariable(modelId string, list []*model.VariableBase) (err error) {
	srv := commonService.NewService(&service.DataModel{})
	srv.OpenTx()
	defer func() {
		if err != nil {
			global.Log.Infof("事务回滚：%s", err.Error())
			srv.Rollback()
		} else {
			global.Log.Infof("事务提交")
			srv.Commit()
		}
	}()

	for _, data := range list {
		err = srv.CreateOrModify(modelId, data)
		if err != nil {
			return
		}
	}

	return
}

// 获取变量列表
func (lc *DataModel) GetVariableList(md_id string) ([]*model.Variable, error) {
	srv := commonService.NewService(&service.DataModel{})
	return srv.GetVarsByKey(md_id)
}

// 删除变量
func (lc *DataModel) DeleteVariable(md_id, id string) error {
	srv := commonService.NewService(&service.DataModel{})
	return srv.DelVarByKey(md_id, id)
}
