package logic

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/issueye/code_magic/backend/code_engine"
	"github.com/issueye/code_magic/backend/common/model"
	commonService "github.com/issueye/code_magic/backend/common/service"
	"github.com/issueye/code_magic/backend/global"
	"github.com/issueye/code_magic/backend/service"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type MainFunc func() (string, error)

const CONSOLE_EVENT = "console"

type CodeLogic struct {
	ctx context.Context
}

func NewCodeLogic(ctx context.Context) *CodeLogic {
	return &CodeLogic{
		ctx: ctx,
	}
}

// 运行代码
func (lc *CodeLogic) RunCode(dmId string, isTest bool, tpCodeId string) error {
	// 结束所有事件
	defer func() {
		runtime.EventsOffAll(lc.ctx)
	}()

	global.Log.Infof("开始运行代码，dmId: %s", dmId)

	// 查询模板信息
	runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("开始获取模板[%s]信息", dmId))
	srv := commonService.NewService(&service.DataModel{})
	info, err := srv.GetById(dmId)
	if err != nil {
		runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("获取模板[%s]失败 %s", dmId, err))
		return err
	}

	// 获取模板明细信息
	runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("开始获取模板[%s]明细信息", dmId))
	columns, err := srv.GetModelInfo(dmId)
	if err != nil {
		runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("获取模板明细失败 %s", err))
		return err
	}

	// 将信息注入到虚拟机中
	runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("开始注入模板[%s]信息", dmId))
	logPath := filepath.Join("runtime", "logs")
	core := code_engine.NewCore(code_engine.OptionLog(logPath, global.Logger.Named("code_engine")))
	globalPath := filepath.Join("runtime", "static", "code")
	core.SetGlobalPath(globalPath)
	core.SetProperty("dm", "title", info.Title)
	core.SetProperty("dm", "tableName", info.TBName)
	core.SetProperty("dm", "columns", columns)
	core.ConsoleCallBack = func(args ...any) {
		// 去除第一个参数，第一个参数是日志级别
		runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, args[1])
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

			err = lc.runCode(core, tp)
			if err != nil {
				errArr = append(errArr, err)
			}
		}

		if len(errArr) > 0 {
			runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("执行代码失败 %s", errArr))
			return fmt.Errorf("执行代码失败 %s", errArr)
		}
	} else {
		runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("开始执行测试代码[%s]", tpCodeId))
		tp, err := tpSrv.Get(tpCodeId)
		if err != nil {
			runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("获取模板[%s]失败 %s", tpCodeId, err))
			return err
		}

		err = lc.runCode(core, tp)
		if err != nil {
			return err
		}
	}

	return nil
}

// checkMainFunc
// 检查代码中是否包含 main 函数
func (lc *CodeLogic) checkMainFunc(path string, tp *model.CodeTemplate) error {
	// 加载代码，使用正则表达式判断代码中是否包含 main 函数
	runtimePath := filepath.Join("runtime", "static", "code")
	codeData, err := os.ReadFile(filepath.Join(runtimePath, path))
	if err != nil {
		runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("读取代码[%s]失败 %s", tp.Title, err))
		return err
	}

	codeStr := string(codeData)
	// 使用正则表达式判断代码中是否包含 main 函数
	runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("开始匹配代码[%s]中是否包含 main 函数", tp.Title))
	matched, err := regexp.MatchString(`(?s)(?:^|\n)\s*(?://.*?\n|/\*.*?\*/\s*)*\s*func\s+main\s*\(`, codeStr)
	if err != nil {
		runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("正则表达式匹配失败 %s", err))
		return err
	}

	if !matched {
		runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("代码[%s]中未找到 main 函数", tp.Title))
		return fmt.Errorf("代码[%s]中未找到 main 函数", tp.Title)
	}

	return nil
}

func (lc *CodeLogic) runCode(core *code_engine.Core, tp *model.CodeTemplate) (err error) {

	path := fmt.Sprintf("%s.js", tp.FileName)
	global.Log.Infof("开始执行代码 %s", path)

	err = lc.checkMainFunc(path, tp)
	if err != nil {
		return err
	}

	rts := core.GetRts()
	var fn MainFunc
	runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, "开始导出[main]函数")
	err = core.ExportFunc("main", &fn, path, rts)
	if err != nil {
		global.Log.Errorf("导出[main]函数失败 %s", err)
		return err
	}

	if fn != nil {
		code, err := fn()
		if err != nil {
			global.Log.Errorf("执行代码[%s]失败 %s", tp.Title, err)
			runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("执行代码[%s]失败 %s", tp.Title, err))
			return err
		}

		fmt.Printf("执行代码[%s] 结果: %s\n", tp.Title, code)
		runtime.EventsEmit(lc.ctx, CONSOLE_EVENT, fmt.Sprintf("执行代码[%s] 结果: %s\n", tp.Title, code))
	}

	return
}
