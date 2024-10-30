package entity

type (
	Account struct {
		EthereunAddress string
		GastokenBalance float64
		WalletNonce     int32
	}

	ERC20Balance struct {
		EthereunAddress string
		Balance         float64
	}
)
