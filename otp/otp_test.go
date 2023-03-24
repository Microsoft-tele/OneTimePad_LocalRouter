package Otp

import (
	"fmt"
	"testing"
)

func TestOtp_AddOtp(t *testing.T) {
	otp := Otp{}
	otp.InitMysql()
	otp.AddOtp("www.baidu.com", "百度", "18697450302", "Lwj20020302", "1784929126@qq.com")
}
func TestOtp_SelectOtp(t *testing.T) {
	otp := Otp{}
	otp.InitMysql()
	selectOtp, err := otp.SelectOtp("1784929126@qq.com")
	if err != nil {
		fmt.Println("查询失败:", err)
		return
	}
	for i, v := range selectOtp {
		fmt.Printf("[%v : %v]\n", i, v)
	}
}
