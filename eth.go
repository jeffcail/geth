package geth

func (e *EthClient) EthAccount() (string, string) {
	return genEthAddress()
}
