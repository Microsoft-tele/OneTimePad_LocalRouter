package RSAUtils

import (
	"fmt"
	"testing"
)

func TestGenerateRSAKey(t *testing.T) {
	GenerateRSAKey(1024)
}

func TestRSA_Encrypt(t *testing.T) {

	GenerateRSAKey(2048)
	res_byte := RSA_Encrypt([]byte("Lwj20020302"), ".\\rsa\\keys\\pub.pem")
	fmt.Println("Res:", string(res_byte))
	dec_bytes := RSA_Decrypt(res_byte, ".\\rsa\\keys\\pri.pem")
	fmt.Println("Dec:", string(dec_bytes))
	
}
