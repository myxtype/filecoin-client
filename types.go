package filecoin

import (
	"github.com/ipfs/go-cid"
	"github.com/shopspring/decimal"
)

type KeyInfo struct {
	Type       string `json:"Type"` // secp256k1
	PrivateKey []byte `json:"PrivateKey"`
}

type Signature struct {
	Type byte   `json:"Type"`
	Data []byte `json:"Data"`
}

type Message struct {
	Version    uint64          `json:"Version"`
	To         string          `json:"To"`
	From       string          `json:"From"`
	Nonce      uint64          `json:"Nonce"`
	Value      decimal.Decimal `json:"Value"`
	GasLimit   int64           `json:"GasLimit"`
	GasFeeCap  decimal.Decimal `json:"GasFeeCap"`
	GasPremium decimal.Decimal `json:"GasPremium"`
	Method     uint64          `json:"Method"`
	Params     []byte          `json:"Params"`
}

type MessageSendSpec struct {
	MaxFee string `json:"MaxFee"`
}

type SignedMessage struct {
	Message   *Message   `json:"Message"`
	Signature *Signature `json:"Signature"`
}

type BlockMessages struct {
	BlsMessages   []*Message       `json:"BlsMessages"`
	SecpkMessages []*SignedMessage `json:"SecpkMessages"`
	Cids          []*cid.Cid       `json:"Cids"`
}

type Actor struct {
	Code    cid.Cid         `json:"Code"`
	Head    cid.Cid         `json:"Head"`
	Nonce   uint64          `json:"Nonce"`
	Balance decimal.Decimal `json:"Balance"`
}

type Ticket struct {
	VRFProof []byte
}

type BlockHeader struct {
	Miner                 string
	Ticket                *Ticket
	Parents               []cid.Cid
	ParentWeight          decimal.Decimal
	Height                int64
	ParentStateRoot       cid.Cid
	ParentMessageReceipts cid.Cid
	Messages              cid.Cid
	BLSAggregate          *Signature
	Timestamp             uint64
	BlockSig              *Signature
	ForkSignaling         uint64
	ParentBaseFee         decimal.Decimal
	validated             bool
}

type TipSet struct {
	Cids   []cid.Cid
	Blocks []*BlockHeader
	Height int64
}

type MessageReceipt struct {
	// Ok = ExitCode(0)
	ExitCode int64 // 状态为0表示成功
	Return   []byte
	GasUsed  int64
}
