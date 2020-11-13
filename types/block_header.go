package types

import (
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/ipfs/go-cid"
	"github.com/shopspring/decimal"
)

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
