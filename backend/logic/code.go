package logic

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/issueye/code_magic/backend/code_engine"
	"github.com/issueye/code_magic/backend/common/model"
	commonService "github.com/issueye/code_magic/backend/common/service"
	"github.com/issueye/code_magic/backend/global"
	"github.com/issueye/code_magic/backend/service"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type MainFunc func() (string, error)

// 运行代码
func RunCode(ctx context.Context, dmId string, isTest bool, tpCodeId string) error {
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
	core.SetProperty("dm", "tableName", info.TBName)
	core.SetProperty("dm", "columns", columns)
	core.ConsoleCallBack = func(args ...any) {
		// 去除第一个参数，第一个参数是日志级别
		runtime.EventsEmit(ctx, "console", args[1])
	}

	// 获取代码模板的脚本
	tpSrv := commonService.NewService(&service.Template{})

	if !isTest {
		tps, err := tpSrv.Gets(info.TPIds)
		if err != nil {
			return err
		}

		errArr := make([]error, 0)
		for _, tp := range tps {

			err = runCode(ctx, core, tp)
			if err != nil {
				errArr = append(errArr, err)
			}
		}

		if len(errArr) > 0 {
			return fmt.Errorf("执行代码失败 %s", errArr)
		}
	} else {
		tp, err := tpSrv.Get(tpCodeId)
		if err != nil {
			return err
		}

		err = runCode(ctx, core, tp)
		if err != nil {
			return err
		}
	}

	return nil
}

func runCode(ctx context.Context, core *code_engine.Core, tp *model.CodeTemplate) (err error) {
	rts := core.GetRts()
	var fn MainFunc
	path := fmt.Sprintf("%s.js", tp.FileName)
	global.Log.Infof("开始执行代码 %s", path)
	err = core.ExportFunc("main", &fn, path, rts)
	if err != nil {
		global.Log.Errorf("导出[main]函数失败 %s", err)
		return err
	}

	if fn != nil {
		code, err := fn()
		if err != nil {
			global.Log.Errorf("执行代码[%s]失败 %s", tp.Title, err)
			return err
		}

		fmt.Printf("执行代码[%s] 结果: %s\n", tp.Title, code)

		runtime.EventsEmit(ctx, "console", fmt.Sprintf("执行代码[%s] 结果: %s\n", tp.Title, code))
	}

	return
}
