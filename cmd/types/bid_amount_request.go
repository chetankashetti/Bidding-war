package types

import "math/big"

type BidRequest struct {
	Amount        big.Int `json:"amount"`
	WalletAddress string  `json:"wallet_address"`
}
