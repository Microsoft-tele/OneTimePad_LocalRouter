package main

import (
	"OneTimePadLocalRouter/CryptoUtils"
	"OneTimePadLocalRouter/RSAUtils"
	"OneTimePadLocalRouter/config"
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/experiment", Experiment)

	err := http.ListenAndServe(":8080", nil)
	fmt.Println("Start server:")
	fmt.Println("-------------------------------------------------------------------------------------------")
	if err != nil {
		fmt.Println("监听错误:", err)
		return
	}
}

func Experiment(w http.ResponseWriter, r *http.Request) {
	// 检查请求方法是否是POST
	if r.Method != http.MethodPost {
		http.Error(w, "只允许POST请求", http.StatusMethodNotAllowed)
		return
	}
	Url := r.PostFormValue("url")
	Intro := r.PostFormValue("intro")
	Uname := r.PostFormValue("uname")
	Pwd := r.PostFormValue("pwd")
	Email := r.PostFormValue("email")
	fmt.Println(Url)
	fmt.Println(Intro)
	fmt.Println(Uname)
	fmt.Println(Pwd)
	fmt.Println(Email)

	otp, rsa := Encrypt(Pwd)
	// 在这里进行其他处理
	// 发送响应
	w.WriteHeader(http.StatusOK)

	retContent := otp + "\n" + rsa
	fmt.Fprintf(w, "成功接收并处理请求:"+retContent)
}

func Encrypt(message string) (string, string) {
	StartTime := time.Now().Nanosecond()
	fmt.Println("OTP start time:", StartTime)

	EncryptPwd := CryptoUtils.Encode(message)

	StopTime := time.Now().Nanosecond()
	fmt.Println("OTP stop time:", StopTime)
	sub := (StopTime - StartTime) / 1e3
	fmt.Printf("OTP加密耗时: %v (μs)\n ", sub)

	// 在这里计算一下RSA
	StartTimeOfRSA := time.Now().Nanosecond()
	fmt.Println("RSA start time:", StartTimeOfRSA)

	// 不需要每次都创建RSA密钥对
	RSAUtils.GenerateRSAKey(2048)
	EncryptPwdByRsa := RSAUtils.RsaEncrypt([]byte(message), config.MyConfigParams.Paths.RsaPubPath)
	StopTimeOfRSA := time.Now().Nanosecond()
	fmt.Println("RSA stop time:", StopTimeOfRSA)
	subOfRSA := (StopTimeOfRSA - StartTimeOfRSA) / 1e3

	fmt.Printf("RSA加密耗时: %v (μs)\n", subOfRSA)

	fmt.Printf("OTP加密较RSA加密效率提高: %.6f\n", (float64(subOfRSA)-float64(sub))/float64(subOfRSA))

	fmt.Println("EncryptPwd:", EncryptPwd)
	return EncryptPwd, string(EncryptPwdByRsa)
}
