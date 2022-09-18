package common

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)

// 文本邮件模板
var TextEMailTemplate = `您的账号：%v ，密码：%v ，请及时修改。`

// 超文本标记邮件模版
var HTMLEMailTemplate = `
    <div>
        找到您的信息
        <br>
        <br>
        我们强烈建议您尽快修改您的帐号密码来保障您的帐号安全，为保护你的帐户安全，请勿转发这封邮件。
        <br>
        <br>
        帐号：%v
        <br>
        密码：%v
        <br>
        <br>
        本邮箱只作为发出邮件使用，请不要回复本邮件。
        <br>
        <br>
    </div>
`

// 初始化发送邮件的配置
var (
	emailAddrress     = "smtp.qq.com:25"   // 地址和端口
	emailAuthIdentity = ""                 // 验证身份
	emailAuthUsername = "dabolau@qq.com"   // 验证用户名
	emailAuthPassword = "jasshsdcbgktbjcc" // 验证密码
	emailAuthHost     = "smtp.qq.com"      // 验证主机
)

// 发送文本邮件
func SendTextEMail(to string, subject string, text string) (err error) {
	e := email.NewEmail()
	e.From = emailAuthUsername
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(text)
	err = e.Send(emailAddrress, smtp.PlainAuth(emailAuthIdentity, emailAuthUsername, emailAuthPassword, emailAuthHost))
	return err
}

// 发送超文本标记邮件
func SendHTMLEMail(to string, subject string, html string) (err error) {
	e := email.NewEmail()
	e.From = emailAuthUsername
	e.To = []string{to}
	e.Subject = subject
	e.HTML = []byte(html)
	err = e.Send(emailAddrress, smtp.PlainAuth(emailAuthIdentity, emailAuthUsername, emailAuthPassword, emailAuthHost))
	return err
}
