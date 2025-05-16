package main

import (
	"log"
	"order/internal/database"
	"order/internal/middleware"
	"order/internal/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 加载配置
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// 初始化数据库
	if err := database.Init(viper.GetString("database.path")); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	println("运行成功")

	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(viper.GetString("server.mode"))

	r := gin.Default()

	// 注册路由
	router.RegisterRoutes(r)
	r.Use(middleware.Cors())

	port := viper.GetString("server.port")

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	// r.Run(":3000")

}
