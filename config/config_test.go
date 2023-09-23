package config

import (
	"fmt"
	"testing"
)

func TestParams_InitParams(t *testing.T) {
	ConfigParams, err := InitParams()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ConfigParams.Paths.OtpPath)
}
