package main

import (
	"github.com/olongfen/gorm-gin-admin/src/models"
	"github.com/olongfen/gorm-gin-admin/src/pkg/gredis"
	"github.com/olongfen/gorm-gin-admin/src/pkg/setting"
	"github.com/olongfen/gorm-gin-admin/src/router"
)


func init() {
	// 初始化配置文件
	setting.InitConfig()
	// 初始化模型
	models.InitModel()
	// 初始化redis
	gredis.InitRedisInstance()
	// 初始化路由
	router.InitRouter()
}

func main() {
}
