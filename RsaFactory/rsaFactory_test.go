package RsaFactory

import (
	"fmt"
	"testing"
)

func TestRSAFactory(t *testing.T) {
	tools, err := NewRSAFactory(2048)
	if err != nil {
		fmt.Println("err:", err)
	}
	key, err := tools.EncryptWithNewKey([]byte("Lwj20020302"))
	if err != nil {
		return
	}
	fmt.Println(key)
}
