package main

import (
	"embed"
	"log"
	"runtime"

	"github.com/issueye/code_magic/backend/initialize"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// 初始化程序
	initialize.Initialize()

	// 创建一个 app 对象
	app := NewApp()

	// 创建一个wails app
	err := wails.Run(&options.App{
		Title:     "code_magic",
		Width:     1024,
		Height:    768,
		MinWidth:  1024,
		MinHeight: 768,
		// MaxWidth:          1280,
		// MaxHeight:         800,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         runtime.GOOS != "darwin",
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		AssetServer: &assetserver.Options{
			Assets:     assets,
			Handler:    nil,
			Middleware: nil,
		},
		Menu:             nil,
		Logger:           nil,
		LogLevel:         logger.WARNING,
		OnStartup:        app.startup,
		OnDomReady:       app.domReady,
		OnBeforeClose:    app.beforeClose,
		OnShutdown:       app.shutdown,
		WindowStartState: options.Normal,
		Bind: []interface{}{
			app,
			NewDataModel(),
			NewTemplate(),
		},
		// Windows platform specific options
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			DisableWindowIcon:    true,
			// DisableFramelessWindowDecorations: false,
			WebviewUserDataPath: "",
			BackdropType:        windows.Acrylic,
		},
		// Mac platform specific options
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHidden(),
			Appearance:           mac.NSAppearanceNameVibrantLight,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "code_magic",
				Message: "",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
