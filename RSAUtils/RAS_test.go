package RSAUtils

import (
	"OneTimePadLocalRouter/OtpUtils"
	"fmt"
	"testing"
)

func TestRSA_Encrypt(t *testing.T) {

	GenerateRSAKey(2048)
	resByte := RsaEncrypt([]byte("Lwj20020302"), OtpUtils.ConfigParams.Paths.RsaPubPath)
	fmt.Println("Res:", string(resByte))
	decBytes := RsaDecrypt(resByte, OtpUtils.ConfigParams.Paths.RsaPriPath)
	fmt.Println("Dec:", string(decBytes))

}
