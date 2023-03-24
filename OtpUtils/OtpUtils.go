package OtpUtils

import (
	"OneTimePadLocalRouter/CharsetUtil"
	"OneTimePadLocalRouter/ShellUtils"
	"fmt"
)

func CreateOtp() {
	out := ShellUtils.GetOutFromStdout("quickrdrand_windows_Vw2.exe -k 1024 > ..\\otp\\otp.txt")
	for i, v := range out {
		fmt.Printf("[%v : %v]\n", i, CharsetUtil.CoverGBKToUTF8(v))
	}
	fmt.Println("创建成功:")
}
