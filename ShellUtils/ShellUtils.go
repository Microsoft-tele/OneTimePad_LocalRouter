package ShellUtils

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func GetOutFromStdout(command string) (out []string) {
	shell := exec.Command("cmd.exe", "/C", command)
	//shell.Stdout = os.Stdout

	stdout, err := shell.StdoutPipe()
	if err != nil {
		fmt.Println("Stdout = shell.StdoutPipe err", err)
	}
	reader := bufio.NewReader(stdout)
	err = shell.Start()
	if err != nil {
		fmt.Println("Shell run err:", err)
		return
	}

	out = make([]string, 0)

	for {
		buf, err := reader.ReadString('\n')
		be, _, _ := strings.Cut(buf, "\n")
		if err == io.EOF {
			//fmt.Println("Read finish")
			break
		} else if err != nil {
			fmt.Println("Read err:", err)
		}
		//fmt.Printf(buf)
		out = append(out, be)
	}
	err2 := shell.Wait()
	if err2 != nil {
		return nil
	}
	return
}
