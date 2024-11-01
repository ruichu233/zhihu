package logic

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/gomail.v2"
	"time"
	"zhihu/app/user/internal/svc"
	"zhihu/app/user/pb/user"
)

type SendVerifyCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendVerifyCodeLogic {
	return &SendVerifyCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendVerifyCodeLogic) SendVerifyCode(in *user.SendVerifyCodeRequest) (*user.SendVerifyCodeResponse, error) {
	code, err := generateCaptcha(6)
	if err != nil {
		fmt.Printf("generateCaptcha failed: %v\n", err)
		return nil, err
	}
	if err := sendEmail(in.Email, code); err != nil {
		fmt.Printf("sendEmail failed: %v\n", err)
		return nil, err
	}
	result, err := l.svcCtx.RDB.SetNX(l.ctx, in.Email, code, 10*time.Minute).Result()
	if err != nil {
		return nil, err
	}
	if !result {
		return nil, fmt.Errorf("验证码已存在")
	}
	return &user.SendVerifyCodeResponse{
		Code: code,
	}, nil
}

// 生成随机验证码
func generateCaptcha(length int) (string, error) {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	for i := range bytes {
		bytes[i] = chars[bytes[i]%byte(len(chars))]
	}
	return string(bytes), nil
}

// 发送邮件
func sendEmail(toEmail string, code string) error {
	// 创建邮件消息
	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail)                 // 发件人
	m.SetHeader("To", toEmail)                       // 收件人
	m.SetHeader("Subject", "Your Verification Code") // 邮件主题
	m.SetBody("text/html", fmt.Sprintf(`
    <div style="font-family: Arial, sans-serif; color: #333;">
        <h2 style="color: #4CAF50;">Verification Code</h2>
        <p>Dear User,</p>
        <p>Your verification code is:</p>
        <p style="font-size: 24px; font-weight: bold; color: #4CAF50;">%s</p>
        <p>The code is valid for <strong>10 minutes</strong>.</p>
        <hr style="border: none; border-top: 1px solid #eee;" />
        <p style="font-size: 12px; color: #888;">If you did not request this code, please ignore this email.</p>
    </div>
`, code)) // 邮件正文

	// 使用 SMTP 服务器信息，注意使用 SSL 端口 465
	d := gomail.NewDialer("smtp.qq.com", 465, senderEmail, senderPassword)

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

// QQ 邮箱 SMTP 配置
const (
	smtpHost       = "smtp.qq.com"
	smtpPort       = "587"               // 或者使用 "465"（SSL 端口），但代码稍有不同
	senderEmail    = "3032860034@qq.com" // 你的 QQ 邮箱
	senderPassword = "rnapjqnnlatkdcfa"  // QQ 邮箱授权码（不是密码）
)
