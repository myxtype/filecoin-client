package types

import (
	"github.com/filecoin-project/go-address"
	"github.com/shopspring/decimal"
)

type Message struct {
	Version    uint64          `json:"Version"`
	To         address.Address `json:"To"`
	From       address.Address `json:"From"`
	Nonce      uint64          `json:"Nonce"`
	Value      decimal.Decimal `json:"Value"`
	GasLimit   int64           `json:"GasLimit"`
	GasFeeCap  decimal.Decimal `json:"GasFeeCap"`
	GasPremium decimal.Decimal `json:"GasPremium"`
	Method     uint64          `json:"Method"`
	Params     []byte          `json:"Params"`
}
