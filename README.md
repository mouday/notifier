# Notifier

基于Golang的邮件通知服务

用到的第三方库

- https://github.com/cosmtrek/air
- https://github.com/subosito/gotenv
- https://github.com/jordan-wright/email
- https://github.com/gin-gonic/gin

## 使用方式

1、下载适合所在运行平台的二进制文件

2、配置邮件验证信息`.env`

```bash
# == 应用配置 ==
# 运行模式
GIN_MODE=release
# 监听端口
APP_RUN_ADDRESS=127.0.0.1:8000

# == 邮件配置 ==
# 账号
APP_EMAIL_USERNAME=admin@163.com
# 密码
APP_EMAIL_PASSWORD=123456
# 邮件服务地址
APP_EMAIL_HOST=smtp.163.com
# 邮件服务端口号
APP_EMAIL_PORT=25
```

3、启动服务

```bash
# macos: 
./notifier

# linux: 
./notifier

# windows: 
notifier.exe
```

4、调用api接口发送邮件

发送示例

```json
POST http://127.0.0.1:8001/sendEmail

{
	"to": ["123456@qq.com"],
	"subject": "测试邮件标题",
	"body": "测试邮件内容"
}
```

返回响应

```json
{
	"code": 0,
	"msg": "success"
}
```