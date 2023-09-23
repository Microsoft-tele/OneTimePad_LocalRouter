package FileUtils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func ReadFileContent(filepath string) (list []string) {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)
	if err != nil {
		return
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		line = line[:len(line)-1]
		list = append(list, line)
	}
}

func GetProjectPath() string {
	// 获取当前可执行文件的路径
	executable, err := os.Executable()
	if err != nil {
		fmt.Println("无法获取可执行文件路径：", err)
	}

	// 获取可执行文件所在目录的绝对路径
	executableDir := filepath.Dir(executable)

	// 获取项目的根地址
	projectRoot := filepath.Dir(executableDir)

	return projectRoot
}
