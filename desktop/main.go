package main

import (
	"embed"
	_ "embed"
	"log"
	"log/slog"
	"path/filepath"
	"time"

	"github.com/atticus6/freeTunnel/desktop/config"
	"github.com/atticus6/freeTunnel/desktop/database"
	"github.com/atticus6/freeTunnel/desktop/services"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

func init() {

	application.RegisterEvent[string]("time")
}

func main() {

	dbPath := filepath.Join(config.StoreDir, "db.db")

	logger.Info("应用启动，数据库路径: %s", dbPath)

	if err := database.Init(dbPath); err != nil {
		logger.Fatal("数据库初始化失败: %v", err)
	}
	logger.Info("数据库初始化成功")

	tunnelService := services.NewTunnelService()

	if err := tunnelService.ClearAllQuickTunnelUrls(); err != nil {
		logger.Error(err)
	}

	app := application.New(application.Options{
		Name:        "react-swc-ts",
		Description: "A demo of using raw HTML & CSS",
		LogLevel:    slog.LevelError,
		Services: []application.Service{
			application.NewService(tunnelService),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title: "Window 1",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
	})

	// Create a goroutine that emits an event containing the current time every second.
	// The frontend can listen to this event and update the UI accordingly.
	go func() {
		for {
			now := time.Now().Format(time.RFC1123)
			app.Event.Emit("time", now)
			time.Sleep(time.Second)
		}
	}()

	// Run the application. This blocks until the application has been exited.
	err := app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}
