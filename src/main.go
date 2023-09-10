package main

import (
	"os"

	"github.com/mouday/notifier/src/dto"
	"github.com/mouday/notifier/src/service"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	os.Getenv("APP_ID")

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
		appRunAddress = ":8080"
	}

	server.Run(appRunAddress)
}
