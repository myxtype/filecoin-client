package types

import (
	"github.com/ipfs/go-cid"
	"github.com/myxtype/filecoin-client/crypto"
	"github.com/shopspring/decimal"
	"time"
)

type TipSetKey []cid.Cid

type Version struct {
	Version    string
	APIVersion uint32
	BlockDelay uint64
}

type BeaconEntry struct {
	Round uint64
	Data  []byte
}

type IpldObject struct {
	Cid cid.Cid
	Obj interface{}
}

type HeadChange struct {
	Type string
	Val  *TipSet
}

type ObjStat struct {
	Size  uint64
	Links uint64
}

type KeyInfo struct {
	Type       string `json:"Type"` // secp256k1
	PrivateKey []byte `json:"PrivateKey"`
}

type MessageSendSpec struct {
	MaxFee decimal.Decimal `json:"MaxFee"`
}

type SignedMessage struct {
	Message   *Message          `json:"Message"`
	Signature *crypto.Signature `json:"Signature"`
}

type BlockMessages struct {
	BlsMessages   []*Message       `json:"BlsMessages"`
	SecpkMessages []*SignedMessage `json:"SecpkMessages"`
	Cids          []cid.Cid        `json:"Cids"`
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
	BLSAggregate          *crypto.Signature
	Timestamp             uint64
	BlockSig              *crypto.Signature
	ForkSignaling         uint64
	ParentBaseFee         decimal.Decimal
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

type Loc struct {
	File     string
	Line     int
	Function string
}

type GasTrace struct {
	Name string

	Location          []Loc `json:"loc"`
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`
	StorageGas        int64 `json:"sg"`
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`
	VirtualStorageGas int64 `json:"vsg"`

	TimeTaken time.Duration `json:"tt"`
	Extra     interface{}   `json:"ex,omitempty"`
}

type ExecutionTrace struct {
	Msg        *Message
	MsgRct     *MessageReceipt
	Error      string
	Duration   time.Duration
	GasCharges []*GasTrace

	Subcalls []ExecutionTrace
}

type InvocResult struct {
	Msg            *Message
	MsgRct         *MessageReceipt
	ExecutionTrace ExecutionTrace
	Error          string
	Duration       time.Duration
}

type MsgLookup struct {
	Message   cid.Cid // Can be different than requested, in case it was replaced, but only gas values changed
	Receipt   MessageReceipt
	ReturnDec interface{}
	TipSet    TipSetKey
	Height    int64
}
