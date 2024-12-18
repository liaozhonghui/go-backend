package main

import (
	"ai-backend/config"
	"ai-backend/routes"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// 加载环境变量
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 连接数据库
	config.ConnectToDatabase()

	// 设置路由
	r := routes.SetupRouter()

	// 启动服务
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // 默认端口
	}

	err = r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
