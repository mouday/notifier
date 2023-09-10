package main

import (
	"fmt"
	"os"

	"github.com/mouday/notifier/src/dto"
	"github.com/mouday/notifier/src/service"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

const VERSION = "v1.0.3"

func init() {
	gotenv.Load()
}

func main() {
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.ReleaseMode
	}

	// fmt.Println("ginMode:", ginMode)

	gin.SetMode(ginMode)

	// 创建一个服务
	server := gin.Default()

	// 处理请求
	server.POST("/sendEmail", func(ctx *gin.Context) {
		var data dto.SendEmailDTO
		ctx.Bind(&data)

		data.Username = os.Getenv("APP_EMAIL_USERNAME")
		data.Password = os.Getenv("APP_EMAIL_PASSWORD")
		data.Host = os.Getenv("APP_EMAIL_HOST")
		data.Port = os.Getenv("APP_EMAIL_PORT")

		err := service.SendEmail(data)

		// 返回json数据
		if err != nil {
			ctx.JSON(200, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
		} else {
			// 返回json数据
			ctx.JSON(200, gin.H{
				"code": 0,
				"msg":  "success",
			})
		}

	})

	// 启动服务
	appRunAddress := os.Getenv("APP_RUN_ADDRESS")

	if appRunAddress == "" {
		appRunAddress = "127.0.0.1:8000"
	}

	fmt.Println("Notifier version: ", VERSION)
	fmt.Println("run address: ", "http://"+appRunAddress)
	server.Run(appRunAddress)
}
