package models

import (
	"fmt"
	"os"
	"time"

	"github.com/astaxie/beego"
	gomail "gopkg.in/gomail.v2"

	"github.com/go-mailer/validate"
)

// EmailAPI .
type EmailAPI interface {
	generateEmailAuthToken(email string) string
	CheckEmailURL(token string) (bool, string)
	GenerateAuthURL(email string) string
	InitSendCfg(recAddr, recName string) error
	SetTheme(theme string)
	SetEmailContent(content string)
	SendEmail() error
}

// Email to send email
type Email struct {
	senderAddr string
	senderName string
	senderkey  string
	recAddr    string
	recName    string
	theme      string
	content    string
	tokenv     *validate.TokenValidate
}

// EmailManager manager emial api
var EmailManager EmailAPI

func init() {
	EmailManager = new(Email)
}

/*
Use "github.com/go-mailer/validate" to check user signup email
*/

// generate auth email token
func (e *Email) generateEmailAuthToken(email string) string {

	// 创建验证信息存储，每10分钟执行一次GC
	store := validate.NewMemoryStore(time.Second * 60 * 60)
	// 创建验证信息管理器，验证信息的过期时间为1小时
	tokenv := validate.NewTokenValidate(store, validate.Config{Expire: time.Second * 60 * 60})
	// 使用邮箱生成验证令牌
	token, err := tokenv.Generate(email)
	if err != nil {
		panic(err)
	}

	e.tokenv = tokenv
	// glog.Infof("check user email token[%s]\n", token)

	return token
}

// CheckEmailURL ..
func (e *Email) CheckEmailURL(token string) (bool, string) {
	// 验证令牌
	isValid, email, err := e.tokenv.Validate(token)
	if err != nil {
		panic(err)
	}
	// glog.Infof("valid email[%s], %t\n", email, isValid)

	return isValid, email
}

// GenerateAuthURL ..
func (e *Email) GenerateAuthURL(email string) string {
	URL := fmt.Sprintf("http://%s:8080/register/active?token=%s", beego.AppConfig.String("httpaddr"), e.generateEmailAuthToken(email))
	// glog.Infof("generate auth url[%s] for user email[%s]", URL, email)
	return URL
}

/*
Use github.com/go-gomail/gomail to send email
*/

// InitSendCfg init send email config
func (e *Email) InitSendCfg(recAddr, recName string) error {
	//完善参数检查
	if recAddr == "" || recName == "" {
		err := fmt.Errorf("emailAddress or username is null")
		// glog.Errorf("init email config error: [%s]", err.Error())
		return err
	}
	e.recAddr = recAddr
	e.recName = recName

	if e.senderAddr = beego.AppConfig.String("sendEmailAddr"); e.senderAddr == "" {
		return fmt.Errorf("sendEmailAddr is null")
	}

	if e.senderName = beego.AppConfig.String("senderName"); e.senderName == "" {
		return fmt.Errorf("senderName is null")
	}

	if e.senderkey = os.Getenv("senderkey"); e.senderkey == "" {
		return fmt.Errorf("senderKey is null")
	}

	// glog.Infof("send email to user[%s] addr[%s] content[%s]\n", e.recName, e.recAddr, e.content)

	return nil
}

// SetTheme set email theme
func (e *Email) SetTheme(theme string) {
	e.theme = theme
}

// SetEmailContent ..
func (e *Email) SetEmailContent(content string) {
	e.content = content
}

// SendEmail ..
func (e *Email) SendEmail() error {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", e.senderAddr, e.senderName) // 发件人
	m.SetHeader("To",                                      // 收件人
		m.FormatAddress(e.recAddr, e.recName),
	)
	m.SetHeader("Subject", e.theme)   // 主题
	m.SetBody("text/html", e.content) // 正文

	d := gomail.NewPlainDialer("smtp.qq.com", 465, e.senderAddr, e.senderkey) // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		err := fmt.Errorf("send email error: [%s]", err.Error())
		// glog.Errorln(err)
		return err
	}

	return nil
}
