package initialize

import (
	"fmt"
	"path/filepath"

	"github.com/issueye/code_magic/backend/common/model"
	commonService "github.com/issueye/code_magic/backend/common/service"
	"github.com/issueye/code_magic/backend/global"
	"github.com/issueye/code_magic/backend/pkg/db"
	"github.com/issueye/code_magic/backend/service"
)

// 初始化其他数据
func InitData() {
	path := filepath.Join("runtime", "data", "data.db")
	commonService.DB = db.InitSqlite(path, global.Logger)

	// 初始化表
	err := commonService.DB.AutoMigrate(
		&model.UserInfo{},      // 用户
		&model.UserGroupInfo{}, // 用户组
		&model.DataModel{},     // 数据模型
	)

	if err != nil {
		panic(fmt.Errorf("初始化表失败 %s", err.Error()))
	}

	err = service.NewUserGroup().CreateAdminNonExistent()
	if err != nil {
		panic("检查管理员用户组信息失败 " + err.Error())
	}

	// 创建 admin 用户
	err = service.NewUser().CreateAdminNonExistent()
	if err != nil {
		panic("初始化数据失败 " + err.Error())
	}
}
