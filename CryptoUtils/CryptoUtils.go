package CryptoUtils

import (
	"OneTimePadLocalRouter/FileUtils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Encode(Password string) string {
	// Old solution: Generate a random number to be a start index
	//rand.Seed(time.Now().UnixNano())
	//start := rand.Intn(20000) // TODO 这里其实有BUG需要修改，需要对取出来的密钥流做模运算

	// New solution: Write the every index of start, when we need it we can get start index from config file.
	//for _, v := range ShellUtils.GetOutFromStdout("dir ..\\") {
	//	fmt.Println(CharsetUtil.CoverGBKToUTF8(v))
	//}

	file, err := os.OpenFile("..\\startIndex", os.O_RDONLY, 0777)
	if err != nil {
		return "Open file err"
	}
	reader := bufio.NewReader(file)
	readString, err := reader.ReadString('\n')
	readString = readString[:len(readString)-1]
	if err != nil {
		return "Read file err"
	}

	//fmt.Println("StartIndex is : " + readString)

	start, err := strconv.Atoi(readString)

	//fmt.Println("Start is:", start)

	EncodePwd := ""

	// 加密的关键步骤
	KeyStream := GetKeyStream(len(Password)*2, start) // TODO start使用随机数生成
	for i := 0; i < len(Password); i++ {
		tmp := int(Password[i]) ^ int(KeyStream[i])
		EncodePwd += string(rune(tmp))
	}
	//fmt.Printf("encode: %v\n", EncodePwd+":"+strconv.Itoa(start))

	nextStart := start + len(Password)*2

	//fmt.Println("Len password*2 = ", len(Password)*2)
	//fmt.Println("Next start:", nextStart)

	openFile, err := os.OpenFile("..\\startIndex", os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		return "Open file err"
	}

	_, err = openFile.WriteString(strconv.Itoa(nextStart) + "\n")
	if err != nil {
		return "Write file err"
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
		KeyStream := GetKeyStream(len(EncodePwd)*2, startInt)
		tmp := int(EncodePwd[i]) ^ int(KeyStream[i])
		DecodePwd += string(rune(tmp))
	}
	return DecodePwd
}

func GetKeyStream(length int, start int) string {
	// 密钥流的起始位置
	KeyStream := ReadOtp()[start : start+length]
	//fmt.Println("KeyStream:", KeyStream)
	return KeyStream
}

func ReadOtp() string {
	out := FileUtils.ReadFileContent("..\\otp\\otp.txt")
	otp := ""
	for _, v := range out {
		otp += v
	}
	str := StringStrip(otp)[2:]
	//fmt.Printf("%v", str)
	return str
}
func StringStrip(input string) string {
	if input == "" {
		return ""
	}
	return strings.Join(strings.Fields(input), "")
}
