package geth

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
)

var (
	GenerateEthAddressError = "Generate eth address failed !!!"
	CanNotType              = "cannot assert type: publicKey is not of type *ecdsa.PublicKey"
)

func genEthAddress() {
	var (
		privateKey *ecdsa.PrivateKey
		err        error
	)
	privateKey, err = crypto.GenerateKey()
	if err != nil {
		log.Fatal(GenerateEthAddressError + "【😭】" + err.Error())
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal(CanNotType)
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))
}
