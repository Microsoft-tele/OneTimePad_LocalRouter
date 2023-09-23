package main

//
//import (
//	"OneTimePadLocalRouter/CharsetUtil"
//	"OneTimePadLocalRouter/CryptoUtils"
//	"OneTimePadLocalRouter/OtpUtils"
//	"OneTimePadLocalRouter/RSAUtils"
//	"OneTimePadLocalRouter/ShellUtils"
//	Otp "OneTimePadLocalRouter/otp"
//	"encoding/json"
//	"fmt"
//	"html/template"
//	"net/http"
//	"net/url"
//	"strings"
//	"time"
//)
//
//type Para struct {
//	OtpObjs []Otp.Otp
//	Email   string
//}
//
//func main() {
//	//http.Handle("/css/img/", http.StripPrefix("/css/img/", http.FileServer(http.Dir("../css/img/"))))
//	//http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("../css"))))
//	//http.Handle("/mod/", http.StripPrefix("/mod/", http.FileServer(http.Dir("../mod"))))
//
//	//http.HandleFunc("/createOtp", OtpUtils.CreateOtp) // TODO 需要加以限制，每个用户只能生成一次
//
//	http.HandleFunc("/", Login)
//
//	http.HandleFunc("/showInfo", ShowInfo)
//
//	http.HandleFunc("/encrypt", Encrypt)
//
//	//http.HandleFunc("/getInfo", GetInfo)
//
//	err := http.ListenAndServe(":8080", nil)
//	if err != nil {
//		fmt.Println("监听错误:", err)
//		return
//	}
//}
//
//func ShowInfo(w http.ResponseWriter, r *http.Request) {
//	var email string
//	err := r.ParseForm()
//	if err != nil {
//		fmt.Println("解析Form表单失败:", err)
//		return
//	}
//	for i, v := range r.PostForm {
//		if i == "email" {
//			email = v[0]
//		}
//	}
//
//	resp, err := http.PostForm("https://404060p9q5.zicp.fun/getInfo",
//		url.Values{
//			"email": {email},
//		})
//	if err != nil {
//		fmt.Println("发送POST请求失败:", err)
//	}
//	buf := make([]byte, 0)
//
//	for {
//		tmp := make([]byte, 1024)
//		read, err := resp.Body.Read(tmp)
//		if err != nil {
//			fmt.Println("读取失败:", err)
//			tmp = tmp[:read]
//			buf = append(buf, tmp...)
//			break
//		}
//		tmp = tmp[:read]
//		buf = append(buf, tmp...)
//	}
//	//fmt.Println("读取到的数据:\n", buf)
//
//	var tmpOtpObjs []Otp.Otp
//
//	err = json.Unmarshal(buf, &tmpOtpObjs)
//	if err != nil {
//		fmt.Println("解析JSON是失败:", err)
//		return
//	}
//	fmt.Println("解析成功")
//	for i := 0; i < len(tmpOtpObjs); i++ {
//		tmpOtpObjs[i].Pwd = CryptoUtils.Decode(tmpOtpObjs[i].Pwd)
//		fmt.Printf("[%v : %v]\n", i, tmpOtpObjs[i].Pwd)
//	}
//
//	ExecutePara := Para{
//		OtpObjs: tmpOtpObjs,
//		Email:   email,
//	}
//
//	files, err := template.ParseFiles("../mod/index.html", "../mod/otp.html")
//	err = files.Execute(w, ExecutePara)
//	if err != nil {
//		fmt.Println("解析模板失败:", err)
//		return
//	}
//}
//
//func Encrypt(w http.ResponseWriter, r *http.Request) {
//	r.PostFormValue("url")
//	//fmt.Println("url_t = ", url_t)
//	err := r.ParseForm()
//	if err != nil {
//		fmt.Println("解析Form表单失败:", err)
//		return
//	}
//	var myUrl string
//	var intro string
//	var uname string
//	var pwd string
//	var email string
//	for i, v := range r.PostForm {
//		fmt.Printf("[%v : %v]\n", i, v)
//		if i == "url" {
//			myUrl = v[0]
//		} else if i == "intro" {
//			intro = v[0]
//		} else if i == "uname" {
//			uname = v[0]
//		} else if i == "pwd" {
//			pwd = v[0]
//		} else if i == "email" {
//			email = v[0]
//		}
//	}
//	fmt.Println("myUrl:", myUrl)
//	fmt.Println("intro:", intro)
//	fmt.Println("uname:", uname)
//	fmt.Println("pwd:", pwd)
//	fmt.Println("email:", email)
//
//	StartTime := time.Now().Nanosecond()
//	fmt.Println("OTP start time:", StartTime)
//	EncryptPwd := CryptoUtils.Encode(pwd)
//	StopTime := time.Now().Nanosecond()
//	fmt.Println("OTP stop time:", StopTime)
//	sub := StopTime - StartTime
//	fmt.Printf("OTP加密耗时: %v\n", sub)
//
//	// 在这里计算一下RSA
//	StartTimeOfRSA := time.Now().Nanosecond()
//	fmt.Println("RSA start time:", StartTimeOfRSA)
//	RSAUtils.GenerateRSAKey(2048)
//	RSAUtils.RSA_Encrypt([]byte(pwd), "..\\RSAUtils\\rsa\\keys\\pub.pem")
//	StopTimeOfRSA := time.Now().Nanosecond()
//	fmt.Println("RSA stop time:", StopTimeOfRSA)
//	subOfRSA := StopTimeOfRSA - StartTimeOfRSA
//
//	fmt.Printf("RSA加密耗时: %v\n", subOfRSA)
//
//	fmt.Printf("OTP加密较RSA加密效率提高: %.6f\n", (float64(subOfRSA)-float64(sub))/float64(subOfRSA))
//
//	fmt.Println("EncryptPwd:", EncryptPwd)
//	_, err = http.PostForm("https://404060p9q5.zicp.fun/add",
//		url.Values{
//			"url":   {myUrl},
//			"intro": {intro},
//			"uname": {uname},
//			"pwd":   {EncryptPwd},
//			"email": {email},
//		})
//	if err != nil {
//		fmt.Println("发送POST请求失败:", err)
//	}
//
//	resp, err := http.PostForm("https://404060p9q5.zicp.fun/getInfo",
//		url.Values{
//			"email": {email},
//		})
//	if err != nil {
//		fmt.Println("发送POST请求失败:", err)
//	}
//	buf := make([]byte, 0)
//
//	for {
//		tmp := make([]byte, 1024)
//		read, err := resp.Body.Read(tmp)
//		if err != nil {
//			fmt.Println("读取失败:", err)
//			tmp = tmp[:read]
//			buf = append(buf, tmp...)
//			break
//		}
//		tmp = tmp[:read]
//		buf = append(buf, tmp...)
//	}
//	//fmt.Println("读取到的数据:\n", buf)
//
//	var tmpOtpObjs []Otp.Otp
//
//	err = json.Unmarshal(buf, &tmpOtpObjs)
//	if err != nil {
//		fmt.Println("解析JSON是失败:", err)
//		return
//	}
//	//fmt.Println("解析成功")
//	for i := 0; i < len(tmpOtpObjs); i++ {
//		tmpOtpObjs[i].Pwd = CryptoUtils.Decode(tmpOtpObjs[i].Pwd)
//		//fmt.Printf("[%v : %v]\n", i, tmpOtpObjs[i].Pwd)
//	}
//
//	ExecutePara := Para{
//		OtpObjs: tmpOtpObjs,
//		Email:   email,
//	}
//
//	files, err := template.ParseFiles("../mod/index.html", "../mod/otp.html")
//	err = files.Execute(w, ExecutePara)
//	if err != nil {
//		fmt.Println("解析模板失败:", err)
//		return
//	}
//}
//
//func Login(w http.ResponseWriter, r *http.Request) {
//	flag := IfCreateOtp()
//	if flag == 0 {
//		OtpUtils.CreateOtp()
//	}
//	w.Header().Set("Location", "https://404060p9q5.zicp.fun")
//	w.WriteHeader(302)
//}
//
//func IfCreateOtp() (flag int) {
//	out := ShellUtils.GetOutFromStdout("dir ..\\otp")
//	flag = 0
//	for i, v := range out {
//		GBKv := CharsetUtil.CoverGBKToUTF8(v)
//		fmt.Printf("[%v : %v]\n", i, GBKv)
//		if strings.Contains(GBKv, "otp.txt") {
//			flag = 1
//			break
//		}
//	}
//	if flag == 0 {
//		fmt.Println("Create opt")
//	} else {
//		fmt.Println("Do not need create otp")
//	}
//	return
//}
