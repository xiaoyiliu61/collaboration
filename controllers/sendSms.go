package controllers

import (
	"collaboration/models"
	"collaboration/utils"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type SentSmsController struct {
	beego.Controller
}

func (s *SentSmsController) Post() {
	fmt.Println("发送验证码")
	var smsLogin models.SmsLogin
	err:=s.ParseForm(&smsLogin)
	if err != nil {
		s.Ctx.WriteString("发送验证码失败，请重试！1")
		fmt.Println(err.Error())
		return
	}
	phone:=smsLogin.Phone
	code:=utils.GenRandCode(6)
	result,err:=utils.SendSms(phone,code,utils.SMS_TLP_REGISTER)
	if err != nil {
		fmt.Println(err.Error())
		s.Ctx.WriteString("发送验证码失败，请重试！2")
		return
	}
	if len(result.BizId)==0 {
		s.Ctx.WriteString(result.Message)
		return
	}
	smsRecord:=models.SmsRecord{
		BizId:     result.BizId,
		Phone:     phone,
		Code:      code,
		Status:    result.Code,
		Message:   result.Message,
		TimeStamp: time.Now().Unix(),
	}
	_,err=smsRecord.SaveSmsRecord()
	if err != nil {
		s.Ctx.WriteString("抱歉，获取验证码失败，请重试！3")
		return
	}
	//保存成功 bizId
	s.Data["Phone"]=smsLogin.Phone
	s.Data["BizId"]=smsRecord.BizId
	//验证码登录
	s.TplName="login_sms_second.html"
}

