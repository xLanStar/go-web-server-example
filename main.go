package main

import (
	"fmt"
	"go-web-server-example/internal/router"
	"go-web-server-example/internal/server"
	"go-web-server-example/internal/services/ai"
	"go-web-server-example/internal/services/auth"
	"go-web-server-example/internal/services/db"
	"log"
	"os"
	"os/signal"

	_ "github.com/joho/godotenv/autoload"
)

var (
	webRouter *router.WebRouter
)

func init() {
	// 設定為正式版
	// gin.SetMode(gin.ReleaseMode)

	auth.Init()
	db.Init()
	ai.Init()

	// 建立路由
	webRouter = router.NewWebRouter(os.Getenv("WEB_FOLDER"), os.Getenv("STORAGE_FOLDER"))
}

func main() {
	// DEBUG:
	// user.CreateUser("Lanstar", "Lanstar", "danny95624268@gmail.com", "aa95624268")
	// user.CreateUser("Lanstar2", "Lanstar2", "danny95624268@gmail.com", "aa95624268")

	// 同步 Web 版本
	watcher, startWatch := webRouter.WatchWeb()
	go startWatch()
	defer watcher.Close()

	var data ai.RequestData = ai.RequestData{
		Prompt:  "RD部門的六月支出是多少",
		History: []string{"a", "b"},
	}
	response, _ := ai.Request(&data)
	fmt.Println("response:", response)

	// 啟用 Web HTTP 服務
	go server.ListenAndServe(":"+os.Getenv("WEB_PORT"), webRouter)

	var quit chan os.Signal = make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("關閉伺服器中...")
}
