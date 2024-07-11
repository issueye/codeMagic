package code_engine

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	goja "github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"go.uber.org/zap"
)

const (
	GoPlugins = "lichee"
)

var (
	LogMap = make(map[string]*ZapLogger) // 日志对象
)

type ZapLogger struct {
	log   *zap.Logger
	Close func()
}

// Core
// goja运行时核心的结构体
type Core struct {
	// 对象池
	pool sync.Pool
	// 全局goja加载目录
	globalPath string
	// 外部添加到内部的内容
	pkg map[string]map[string]any
	// 外部注册的模块
	modules map[string]ModuleFunc
	// 对应文件的编译对象
	proMap map[string]*Code
	// 日志对象
	logger *zap.Logger
	// 日志存放路径
	logPath string
	// 日志输出模式
	logMode LogOutMode
	// 输出回调
	ConsoleCallBack ConsoleCallBack
}

type OptFunc = func(*Core)

func NewCore(opts ...OptFunc) *Core {
	c := new(Core)

	c.pkg = make(map[string]map[string]any)
	c.modules = make(map[string]ModuleFunc)
	c.proMap = make(map[string]*Code)

	for _, opt := range opts {
		opt(c)
	}

	// 注册原生模块
	for Name, moduleFn := range c.modules {
		require.RegisterNativeModule(Name, func(runtime *goja.Runtime, module *goja.Object) {
			m := module.Get("exports").(*goja.Object)
			moduleFn(runtime, m)
		})
	}

	return c
}

func (c *Core) InitPool() {
	c.pool = sync.Pool{
		New: func() interface{} {
			jsVM := NewJsVM(
				c.globalPath,
				c.logger,
				c.ConsoleCallBack,
			)

			// 加载其他模块
			for name, mod := range c.pkg {
				gojaMod := jsVM.vm.NewObject()
				for k, v := range mod {
					gojaMod.Set(k, v)
				}

				// 注册模块
				jsVM.vm.Set(name, gojaMod)
			}

			return jsVM
		},
	}
}

// OptionLog
// 配置日志
func OptionLog(path string, log *zap.Logger) OptFunc {
	return func(core *Core) {
		core.logger = log
		core.logPath = path
	}
}

func (c *Core) GetRuntime() *JsVM {
	vm := c.pool.Get().(*JsVM)
	return vm
}

func (c *Core) PutRuntime(vm *JsVM) {
	c.pool.Put(vm)
}

// SetLogPath
// 设置日志路径
func (c *Core) SetLogPath(path string) {
	c.logPath = path
}

// SetLogOutMode
// 日志输出模式
// debug 输出到控制台和输出到日志文件
// release 只输出到日志文件
func (c *Core) SetLogOutMode(mod LogOutMode) {
	c.logMode = mod
}

func (c *Core) SetGlobalPath(path string) {
	c.globalPath = path
}

// Compile
// 编译代码
func (c *Core) Compile(name string, path string) error {

	var tmpPath string
	if c.globalPath != "" {
		tmpPath = filepath.Join(c.globalPath, path)
	} else {
		tmpPath = path
	}

	// 读取文件
	src, err := os.ReadFile(tmpPath)
	if err != nil {
		return err
	}

	// 编译
	c.compile(name, path, string(src))
	return nil
}

// Compile
// 编译代码
func (c *Core) compile(name string, path string, code string) error {
	// 编译文件
	pro, err := goja.Compile(name, code, false)
	if err != nil {
		return err
	}

	c.proMap[name] = &Code{
		Path:    path,
		Program: pro,
	}

	return nil
}

// RunOnce()
func (c *Core) RunOnce(name string, rt *JsVM) error {
	// 只有存在编译对象时，才运行
	if c.proMap[name] != nil {
		return rt.Run(name, c.proMap[name])
	}

	return nil
}

// CallMain
func (c *Core) CallMain(name string, path string) error {

	// 调用 main
	var main func()
	err := c.ExportFunc(name, path, &main)
	if err != nil {
		return err
	}

	// 运行main方法
	main()

	return nil
}

func (c *Core) ExportFunc(name string, path string, fn any) error {
	// 编译
	err := c.Compile(name, path)
	if err != nil {
		return err
	}

	jsvm := c.GetRuntime()
	vm := jsvm.vm

	// 运行
	err = c.RunOnce(name, jsvm)
	if err != nil {
		return err
	}

	nameFunc := vm.Get(name)

	_, ok := goja.AssertFunction(nameFunc)
	if !ok {
		return fmt.Errorf("%s function not found", name)
	}

	// 导出
	return vm.ExportTo(vm.Get(name), fn)
}

// Run
// 运行脚本 文件
func (c *Core) Run(name, path string) error {
	rt := c.GetRuntime()
	defer c.PutRuntime(rt)

	return c.run(name, path, rt)
}

// RunVM
// 运行脚本 文件
func (c *Core) RunVM(name string, path string) error {
	rt := c.GetRuntime()
	defer c.PutRuntime(rt)

	return c.run(name, path, rt)
}

func (c *Core) run(name string, path string, vm *JsVM) error {
	err := c.Compile(name, path)
	if err != nil {
		return err
	}

	return c.RunOnce(name, vm)
}

// RunString
// 运行脚本 字符串
func (c *Core) RunString(name string, src string) error {
	err := c.compile(name, src, "")
	if err != nil {
		return err
	}

	return c.RunOnce(name, nil)
}

// SetGlobalProperty
// 写入数据到全局对象中
func (c *Core) SetGlobalProperty(key string, value any) {
	c.pkg[GoPlugins][key] = value
}

// SetProperty
// 向模块写入变量或者写入方法
func (c *Core) SetProperty(moduleName, key string, value any) {
	mod, ok := c.pkg[moduleName]
	if !ok {
		c.pkg[moduleName] = make(map[string]any)
		mod = c.pkg[moduleName]
	}
	mod[key] = value
}

// RegisterModule
// 注册模块
func (c *Core) RegisterModule(moduleName string, fn ModuleFunc) {
	c.modules[moduleName] = fn
}
