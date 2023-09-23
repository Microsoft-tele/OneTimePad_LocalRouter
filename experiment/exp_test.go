package experiment

import (
	"OneTimePadLocalRouter/config"
	"fmt"
	"testing"
	"time"
)

func TestEncryptOtpList(t *testing.T) {
	startTime := time.Now()
	EncryptOtpList(config.MyConfigParams.Paths.Dataset1mPath)
	endTime := time.Now()

	elapsedTime := endTime.Sub(startTime)
	t.Logf("EncryptOtpList 执行时间: %s", elapsedTime)
}

func TestEncryptRsaList(t *testing.T) {
	startTime := time.Now()
	EncryptRsaList(config.MyConfigParams.Paths.Dataset1mPath)
	endTime := time.Now()

	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("EncryptRsaList 执行时间: %s", elapsedTime)
}

func TestLog(t *testing.T) {
	t.Logf("This is a log")
}

func TestGetResult(t *testing.T) {
	// 测量 EncryptOtpList 执行时间
	otp, rsa, per := Demo(config.MyConfigParams.Paths.Dataset1mPath)
	otp10, rsa10, per10 := Demo(config.MyConfigParams.Paths.Dataset10mPath)
	otp100, rsa100, per100 := Demo(config.MyConfigParams.Paths.Dataset100mPath)
	otp1000, rsa1000, per1000 := Demo(config.MyConfigParams.Paths.Dataset1000mPath)
	fmt.Printf(",otp,rsa,per\n"+
		"1M,%f,%f,%f\n"+
		"10M%f,%f,%f\n"+
		"100M%f,%f,%f\n"+
		"1000M%f,%f,%f",
		otp, rsa, per, otp10, rsa10, per10, otp100, rsa100, per100, otp1000, rsa1000, per1000)
}

func Demo(filePath string) (float64, float64, float64) {
	startTime := time.Now()
	EncryptOtpList(filePath)
	endTime := time.Now()
	elapsedTimeOtp := endTime.Sub(startTime)

	// 测量 EncryptRsaList 执行时间
	startTime = time.Now()
	EncryptRsaList(filePath)
	endTime = time.Now()
	elapsedTimeRsa := endTime.Sub(startTime)

	// 计算百分比差异
	percentageDifference := 100 * float64(elapsedTimeRsa-elapsedTimeOtp) / float64(elapsedTimeRsa)

	fmt.Printf("EncryptOtpList 执行时间: %s\n", elapsedTimeOtp)
	fmt.Printf("EncryptRsaList 执行时间: %s\n", elapsedTimeRsa)
	fmt.Printf("OTP 和 RSA 加密效率百分比差异: %.2f%%\n", percentageDifference)
	return float64(elapsedTimeOtp), float64(elapsedTimeRsa), percentageDifference
}
