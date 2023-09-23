package FileUtils

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestReadFileContent(t *testing.T) {
	out := ReadFileContent("index.txt")
	fmt.Println(out)
	file, err := os.OpenFile("index.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Err:", err)
		return
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		line = line[:len(line)-1]
		fmt.Println(line)
	}
}

func TestGetProjectPath(t *testing.T) {
	GetProjectPath()
}
