package CryptoUtils

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(Decode(Encode("Lwj20020302;")))
}
