package OtpUtils

import (
	"OneTimePadLocalRouter/CharsetUtil"
	"OneTimePadLocalRouter/ShellUtils"
	"OneTimePadLocalRouter/config"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"syscall"
	"time"
	"unsafe"
)

var (
	user32         = syscall.NewLazyDLL("user32.dll")
	getCursorPos   = user32.NewProc("GetCursorPos")
	screenToClient = user32.NewProc("ScreenToClient")
)

func CreateOtp_() {
	out := ShellUtils.GetOutFromStdout("quickrdrand_windows_Vw2.exe -k 1024 > ..\\otp\\otp.txt")
	for i, v := range out {
		fmt.Printf("[%v : %v]\n", i, CharsetUtil.CoverGBKToUTF8(v))
	}
	fmt.Println("创建成功:")
}

type POINT struct {
	X, Y int32
}

func GetPos() POINT {
	var pt POINT
	_, _, err := getCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	if err != nil {
		fmt.Println("log:", err)
	}
	fmt.Printf("鼠标坐标：X：%d, Y：%d\n", pt.X, pt.Y)
	return pt
}

func ComplexMath(seedX, seedY int64) int64 {
	// 进行一些复杂的数学处理，例如位运算、模运算、指数运算等
	result := (seedX ^ seedY) + int64(math.Pow(float64(seedX), 2)) - int64(math.Sqrt(float64(seedY)))
	result = result % 1000000 // 取模运算，将结果限制在一定范围内
	return result
}

func GenerateRandomSeed() int64 {
	// 使用当前时间戳作为随机数种子的一部分
	timestamp := time.Now().UnixNano()

	// 获取鼠标坐标
	mousePos := GetPos()

	// 将鼠标坐标和时间戳进行数学变换
	seedX := (timestamp + int64(mousePos.X)) % 1000000
	seedY := (timestamp + int64(mousePos.Y)) % 1000000

	// 进行一些数学处理以确保均匀分布
	seed := ComplexMath(seedX, seedY)

	return seed
}

func GenerateOTP(seed int64, length int) string {
	// 使用生成的种子初始化随机数生成器
	rand.Seed(seed)

	// 生成随机的 OTP 密钥
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println("生成 OTP 密钥时出错：", err)
		return ""
	}

	// 将 OTP 密钥以十六进制字符串形式返回
	return hex.EncodeToString(key)
}

// CreateOtp generates a One-Time Password (OTP) of the specified length.
// It returns the generated OTP as a string.
//
// Parameters:
//
//	length: the length of the generated OTP
//
// Returns:
//
//	string: the generated OTP
func CreateOtp(length int) {
	// 检查文件是否存在
	if _, err := os.Stat(config.MyConfigParams.Paths.OtpPath); os.IsNotExist(err) {
		// 如果文件不存在，调用生成 OTP 的函数
		file, _ := os.Create(config.MyConfigParams.Paths.OtpPath)
		fileWriter := io.Writer(file)
		otpString := GenerateOTP(GenerateRandomSeed(), length)
		startIndex := "0"
		otpString = startIndex + "@" + otpString
		fp, err := fmt.Fprint(fileWriter, otpString)
		if err != nil {
			fmt.Println("err:", err, fp)
		}
		fmt.Println("OTP 密码本创建成功")
	} else {
		// 如果文件已经存在，打印提示
		fmt.Println("OTP 密码本已经存在")
	}
}

func AppendOtp(length int) {
	file, err := os.OpenFile(config.MyConfigParams.Paths.OtpPath, os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Open file err:", err)
	}
	fileWriter := io.Writer(file)
	otpString := GenerateOTP(GenerateRandomSeed(), length)
	fp, err := fmt.Fprint(fileWriter, otpString)
	if err != nil {
		fmt.Println("err:", err, fp)
	}
	file.Close()
}
