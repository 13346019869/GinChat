package main

import (
	"ginchat/models"
	"ginchat/router"
	"ginchat/utils"
	"github.com/spf13/viper"
	"time"
)

func main() {
	//utils.InitConfig()
	//utils.InitMySQL()
	//
	//db := utils.DB
	//
	////执行迁移文件 生成数据表
	//db.AutoMigrate(&models.UserBasic{})
	//db.AutoMigrate(&models.Message{})
	//db.AutoMigrate(&models.GroupBasic{})
	//db.AutoMigrate(&models.Contact{})
	//db.AutoMigrate(&models.Community{})
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	InitTimer()
	r := router.Router()                  // router.Router()
	r.Run(viper.GetString("port.server")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// 初始化定时器
func InitTimer() {
	utils.Timer(time.Duration(viper.GetInt("timeout.DelayHeartbeat"))*time.Second, time.Duration(viper.GetInt("timeout.HeartbeatHz"))*time.Second, models.CleanConnection, "")
}
