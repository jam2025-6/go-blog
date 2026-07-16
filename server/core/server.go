package core

import (
	"fmt"
	"server/global"
	"server/initialize"
	"server/service"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	addr := global.Config.System.Addr()
	Router := initialize.InitRouter()
	// LoadAll 从数据库加载所有的JWT黑名单并加入缓存
	service.LoadAll()

	s := initServer(addr, Router)

	fmt.Printf("Server is running at http://%s\n", addr)
	global.Log.Info("server run success on ", zap.String("address", addr))
	global.Log.Error(s.ListenAndServe().Error())
}
