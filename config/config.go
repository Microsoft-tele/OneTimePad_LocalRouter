package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var (
	MyConfigParams, _ = InitParams()
)

type Paths struct {
	OtpPath          string `json:"otp_path"`
	StartIndex       string `json:"start_index"`
	RsaPriPath       string `json:"rsa_pri_path"`
	RsaPubPath       string `json:"rsa_pub_path"`
	Dataset1mPath    string `json:"dataset_1M_path"`
	Dataset10mPath   string `json:"dataset_10M_path"`
	Dataset100mPath  string `json:"dataset_100M_path"`
	Dataset1000mPath string `json:"dataset_1000M_path"`
}

type Params struct {
	ProjectPath string `json:"project_path"`
	Paths       Paths  `json:"paths"`
}

func InitParams() (Params, error) {
	p := Params{}

	file, err := os.Open("C:\\Users\\Nahida\\Desktop\\git\\OneTimePad_LocalRouter\\config\\config.json")
	if err != nil {
		fmt.Println("无法打开配置文件:", err)
		return Params{}, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("err:", err)
			return
		}
	}(file)
	// 解码 JSON 配置文件到 Params 结构体
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&p); err != nil {
		fmt.Println("无法解码配置文件:", err)
		fmt.Println("err:", err)
	}

	// Please register all paths at here @TODO remember to register
	p.Paths.OtpPath = filepath.Join(p.ProjectPath, p.Paths.OtpPath)
	p.Paths.StartIndex = filepath.Join(p.ProjectPath, p.Paths.StartIndex)
	p.Paths.RsaPubPath = filepath.Join(p.ProjectPath, p.Paths.RsaPubPath)
	p.Paths.RsaPriPath = filepath.Join(p.ProjectPath, p.Paths.RsaPriPath)

	p.Paths.Dataset1mPath = filepath.Join(p.ProjectPath, p.Paths.Dataset1mPath)
	p.Paths.Dataset10mPath = filepath.Join(p.ProjectPath, p.Paths.Dataset10mPath)
	p.Paths.Dataset100mPath = filepath.Join(p.ProjectPath, p.Paths.Dataset100mPath)
	p.Paths.Dataset1000mPath = filepath.Join(p.ProjectPath, p.Paths.Dataset1000mPath)
	fmt.Println("配置文件已成功初始化")
	return p, nil
}
