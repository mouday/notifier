package dto

type SendEmailDTO struct {
	Username string   `form:"username" binding:"-" json:"username"`
	Password string   `form:"password" binding:"-" json:"password"`
	Host     string   `form:"host" binding:"-" json:"host"`
	Port     string   `form:"port" binding:"-" json:"port"`
	To       []string `form:"to" binding:"required" json:"to"`
	Subject  string   `form:"subject" binding:"required" json:"subject"`
	Body     string   `binding:"required" json:"body"`
}
