package geth

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math"
	"math/big"
	"regexp"
)

var (
	GenerateEthAddressError = "Generate eth address failed !!!"
	CanNotType              = "cannot assert type: publicKey is not of type *ecdsa.PublicKey"
)

// genEthAddress
func genEthAddress() (string, string) {
	var (
		privateKey *ecdsa.PrivateKey
		err        error
	)
	privateKey, err = crypto.GenerateKey()
	if err != nil {
		log.Fatal(GenerateEthAddressError + "【😭】" + err.Error())
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyStr := hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal(CanNotType)
	}

	//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return privateKeyStr, address

	//hash := sha3.NewLegacyKeccak256()
	//hash.Write(publicKeyBytes[1:])
	//encode := hexutil.Encode(hash.Sum(nil)[12:])
	//fmt.Println()
}

// hexToAddress
func hexToAddress(s string) common.Address {
	return common.HexToAddress(s)
}

// balanceToWei
// 以太坊中的数字是使用尽可能小的单位来处理的，因为它们是定点精度，在ETH中它是wei。要读取ETH值，您必须做计算wei/10^18
func balanceToWei(balance *big.Int) *big.Float {
	b := new(big.Float)
	b.SetString(b.String())
	ethValue := new(big.Float).Quo(b, big.NewFloat(math.Pow10(18)))
	return ethValue
}

var (
	regComStr = "^0x[0-9a-fA-F]{40}$"
)

// ValidaEthAddress 校验地址是否有效
func ValidaEthAddress(address string) bool {
	return regexp.MustCompile(regComStr).MatchString(address)
}
