package FileUtils

import (
	"bufio"
	"os"
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
