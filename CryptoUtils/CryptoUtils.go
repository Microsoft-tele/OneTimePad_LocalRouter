package CryptoUtils

import (
	"OneTimePadLocalRouter/OtpUtils"
	"OneTimePadLocalRouter/config"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// Encode this function will encrypt the password through otp
//
// Parameters:
//
// Password: the message will be encoded
//
// Return:
//
// encoded message
func Encode(Password string) string {
	EncodePwd := ""

	// 加密的关键步骤
	KeyStream, start := GetKeyStream(len(Password)*2, true, 0)
	for i := 0; i < len(Password); i++ {
		tmp := int(Password[i]) ^ int(KeyStream[i])
		EncodePwd += string(rune(tmp))
	}
	return EncodePwd + ":" + strconv.Itoa(start)
}

func Decode(EncodePwd string) string {
	DecodePwd := ""
	ValidEncodePwd := strings.Split(EncodePwd, ":")
	Encode := ValidEncodePwd[0]
	start := ValidEncodePwd[1]
	startInt, err := strconv.Atoi(start)
	if err != nil {
		fmt.Println("转换失败:", err)
	}
	//fmt.Println("Encode:", Encode)
	//fmt.Println("start:", start)
	for i := 0; i < len(Encode); i++ {
		KeyStream, _ := GetKeyStream(len(EncodePwd)*2, false, startInt)
		tmp := int(EncodePwd[i]) ^ int(KeyStream[i])
		DecodePwd += string(rune(tmp))
	}
	return DecodePwd
}

func GetKeyStream(length int, isEncrypt bool, startIndex int) (string, int) {
	// 密钥流的起始位置
	var keyStream string
Loop:
	content := ReadOtp()
	split := strings.Split(content, "@")

	if isEncrypt {
		// If startIndex + length >= len(otp pad)
		startIndex, _ = strconv.Atoi(split[0])
		if startIndex+length >= len(split[1]) {
			fmt.Println("Detecting otp has been run out:")
			OtpUtils.AppendOtp(1024)
			goto Loop
		}
		nextStartIndex := startIndex + length
		// 将nextStartIndex写回密码本
		content = strconv.Itoa(nextStartIndex) + "@" + split[1]
		err := WriteOtp(content)
		if err != nil {
			fmt.Println("写入密码本失败:", err)
		}
	}

	keyStream = split[1][startIndex : startIndex+length]
	//fmt.Println("KeyStream:", KeyStream)
	return keyStream, startIndex
}

var otpLock sync.Mutex

func ReadOtp() string {
	otpLock.Lock()
	defer otpLock.Unlock()
	content, err := os.ReadFile(config.MyConfigParams.Paths.OtpPath)
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return ""
	}
	// 将文件内容转换为字符串并返回
	return string(content)
}

func WriteOtp(content string) error {
	otpLock.Lock()
	defer otpLock.Unlock()
	err := os.WriteFile(config.MyConfigParams.Paths.OtpPath, []byte(content), 0666)
	if err != nil {
		fmt.Println("写入文件失败:", err)
		return err
	}
	return nil
}
