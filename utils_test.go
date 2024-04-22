package geth

import (
	"fmt"
	"testing"
)

func TestValidaEthAddress(t *testing.T) {
	address := "0xb45D0616870aAf9AC4970f111592f85107E1C831"
	ok := ValidaEthAddress(address)
	fmt.Println(ok)

	address2 := "0xb45D0616870aAf9AC4970f111592f85107E1C8311212"
	ok2 := ValidaEthAddress(address2)
	fmt.Println(ok2)

	address3 := "0xZafsdXb5d4c32345ced77393b3530b1eed0f346429d"
	ok3 := ValidaEthAddress(address3)
	fmt.Println(ok3)
}
