package experiment

import (
	"OneTimePadLocalRouter/CryptoUtils"
	"OneTimePadLocalRouter/RSAUtils"
	"OneTimePadLocalRouter/config"
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"os"
	"sync"
)

func EncryptOtpList(filePath string) {
	dataset, err := os.Open(filePath)
	defer dataset.Close()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	df := dataframe.ReadCSV(dataset)
	pwdCol := df.Select([]string{"password"})

	// 创建等待组，用于等待所有 goroutine 完成
	var wg sync.WaitGroup

	// 限制同时运行的协程数量
	maxConcurrent := 50 // 指定最大同时运行的协程数量
	sema := make(chan struct{}, maxConcurrent)

	// 并发处理密码加密
	for _, row := range pwdCol.Records() {
		wg.Add(1)
		sema <- struct{}{} // 占用一个信号量，控制同时运行的协程数量
		go func(password string) {
			defer func() {
				<-sema // 释放信号量
				wg.Done()
			}()
			EncodedPwd := CryptoUtils.Encode(password)
			fmt.Println("EncPwd:", EncodedPwd)
		}(row[0])
	}

	// 等待所有 goroutine 完成
	wg.Wait()
}

func EncryptRsaList(filePath string) {
	dataset, err := os.Open(filePath)
	defer dataset.Close()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	df := dataframe.ReadCSV(dataset)
	pwdCol := df.Select([]string{"password"})

	// 创建等待组，用于等待所有 goroutine 完成
	var wg sync.WaitGroup

	// 限制同时运行的协程数量
	maxConcurrent := 50 // 指定最大同时运行的协程数量
	sema := make(chan struct{}, maxConcurrent)

	// 并发处理密码RSA加密
	for _, row := range pwdCol.Records() {
		wg.Add(1)
		sema <- struct{}{} // 占用一个信号量，控制同时运行的协程数量
		go func(password string) {
			defer func() {
				<-sema // 释放信号量
				wg.Done()
			}()
			fmt.Println("Password:", password)
			// RSAUtils.GenerateRSAKey(2048)
			EncryptPwdByRsa := RSAUtils.RsaEncrypt([]byte(password), config.MyConfigParams.Paths.RsaPubPath)
			fmt.Println("EncPwd:", EncryptPwdByRsa)
		}(row[0])
	}

	// 等待所有 goroutine 完成
	wg.Wait()
}
