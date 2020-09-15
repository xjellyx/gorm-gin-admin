package main

import (
	"fmt"
	"github.com/olongfen/gorm-gin-admin/src/models"
	_ "github.com/olongfen/gorm-gin-admin/src/models"
	_ "github.com/olongfen/gorm-gin-admin/src/pkg/gredis"
	"time"
	//_ "github.com/olongfen/gorm-gin-admin/src/router"
	//_ "github.com/olongfen/gorm-gin-admin/src/setting"
)

func init() {
	// 初始化配置文件
	// setting.InitConfig()
	// 初始化模型
	//models.InitModel()
	//// 初始化redis
	//gredis.InitRedisInstance()
	//// 初始化路由
	//router.InitRouter()
}

func main() {
	u := models.UserBase{}
	time.Sleep(time.Second * 3)
	u.GetByUId("a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9")
	fmt.Println(u)
}
