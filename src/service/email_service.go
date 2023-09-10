package service

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"github.com/mouday/notifier/src/dto"

	"github.com/jordan-wright/email"
)

func SendEmail(data dto.SendEmailDTO) error {

	// 简单设置 log 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	username := strings.Split(data.Username, "@")
	em.From = fmt.Sprintf("%s <%s>", username[0], data.Username)
	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = data.To

	// 设置主题
	em.Subject = data.Subject

	// 简单设置文件发送的内容，暂时设置成纯文本
	em.Text = []byte(data.Body)

	//设置服务器相关的配置
	addr := fmt.Sprintf("%s:%s", data.Host, data.Port)
	err := em.Send(addr, smtp.PlainAuth("", data.Username, data.Password, data.Host))

	if err != nil {
		log.Println("send error:", err)
	} else {
		log.Println("send successfully ... ")
	}

	return err

}
