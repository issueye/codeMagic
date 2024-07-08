package logic

import (
	"fmt"
	"path/filepath"

	"github.com/issueye/code_magic/backend/code_engine"
	commonService "github.com/issueye/code_magic/backend/common/service"
	"github.com/issueye/code_magic/backend/global"
	"github.com/issueye/code_magic/backend/service"
)

type MainFunc func() string

// 运行代码
func RunCode(dmId string) error {
	global.Log.Infof("开始运行代码，dmId: %s", dmId)

	// 查询模板信息
	srv := commonService.NewService(&service.DataModel{})
	info, err := srv.GetById(dmId)
	if err != nil {
		return err
	}

	// 获取模板明细信息
	columns, err := srv.GetModelInfo(dmId)
	if err != nil {
		return err
	}

	// 将信息注入到虚拟机中
	logPath := filepath.Join("runtime", "logs")
	core := code_engine.NewCore(code_engine.OptionLog(logPath, global.Logger.Named("code_engine")))
	globalPath := filepath.Join("runtime", "static", "code")
	core.SetGlobalPath(globalPath)
	core.SetProperty("dm", "title", info.Title)
	core.SetProperty("dm", "columns", columns)

	// 获取代码模板的脚本
	tpSrv := commonService.NewService(&service.Template{})
	tps, err := tpSrv.Gets(info.TPIds)
	if err != nil {
		return err
	}

	errArr := make([]error, 0)
	for _, tp := range tps {
		rts := core.GetRts()
		var fn MainFunc
		path := fmt.Sprintf("%s.js", tp.FileName)
		global.Log.Infof("开始执行代码 %s", path)
		err = core.ExportFunc("main", &fn, path, rts)
		if err != nil {
			global.Log.Errorf("导出[main]函数失败 %s", err)
			errArr = append(errArr, err)
			continue
		}

		if fn != nil {
			code := fn()
			fmt.Printf("执行代码[%s] 结果: %s\n", tp.Title, code)
		}
	}

	if len(errArr) > 0 {
		return fmt.Errorf("执行代码失败 %s", errArr)
	}

	return nil
}
