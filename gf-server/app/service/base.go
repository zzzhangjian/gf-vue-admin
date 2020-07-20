package service

import (
	"errors"
	"gf-server/app/model/admins"
	"github.com/gogf/gf/frame/g"
	"github.com/mojocn/base64Captcha"
	uuid "github.com/satori/go.uuid"
)

var store = base64Captcha.DefaultMemStore

// Register 注册
func Register(u *admins.Entity) (err error) {
	if !admins.RecordNotFound(g.Map{"username": u.Username}) {
		return errors.New("用户已存在,注册失败")
	}
	u.Uuid = uuid.NewV4().String()
	if err = u.EncryptedPassword(); err != nil { // 哈希加密
		return errors.New("注册失败")
	}
	if _, err = admins.Insert(u); err != nil {
		return errors.New("注册失败")
	}
	return nil
}

// Captcha Verification code generation
// Captcha 验证码生成
func Captcha() (id string, b64s string, err error) {
	driver := base64Captcha.NewDriverDigit(g.Cfg().GetInt("captcha.ImgHeight"), g.Cfg().GetInt("captcha.ImgWidth"), g.Cfg().GetInt("captcha.KeyLong"), 0.7, 80) 	// 字符,公式,验证码配置, 生成默认数字的driver
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err = cp.Generate()
	return id, b64s, err
}
