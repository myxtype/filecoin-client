package main

import (
	"context"
	"encoding/hex"
	"github.com/filecoin-project/go-address"
	"github.com/myxtype/filecoin-client"
	"github.com/myxtype/filecoin-client/local"
	"github.com/myxtype/filecoin-client/types"
	"github.com/shopspring/decimal"
)

func main() {
	address.CurrentNetwork = address.Mainnet

	ki, addr, err := local.WalletNew(types.KTSecp256k1)
	if err != nil {
		panic(err)
	}

	println(hex.EncodeToString(ki.PrivateKey))
	println(addr.String())

	// 50a5e6234f5fdfc026bd889347409e11b6ef5b6034a7b0572d7e24ed1e9ba0e4
	// f1dynqskhlixt5eswpff3a72ksprqmeompv3pbesy

	to, _ := address.NewFromString("f1yfi4yslez2hz3ori5grvv3xdo3xkibc4v6xjusy")

	msg := &types.Message{
		Version:    0,
		To:         to,
		From:       *addr,
		Nonce:      0,
		Value:      decimal.NewFromInt(10000),
		GasLimit:   0,
		GasFeeCap:  decimal.NewFromInt(10000),
		GasPremium: decimal.NewFromInt(10000),
		Method:     0,
		Params:     nil,
	}

	s, err := local.WalletSignMessage(types.KTSecp256k1, ki.PrivateKey, msg)
	if err != nil {
		panic(err)
	}

	println(hex.EncodeToString(s.Signature.Data))
	// 47bcbb167fd9040bd02dba02789bc7bc0463c290db1be9b07065c12a64fb84dc546bef7aedfba789d0d7ce2c4532f8fa0d2dd998985ad3ec1a8b064c26e4625a01

	client := filecoin.New("https://1lB5G4SmGdSTikOo7l6vYlsktdd:b58884915362a99b4fc18c2bf8af8358@filecoin.infura.io")

	mid, err := client.MpoolPush(context.Background(), s)
	if err != nil {
		panic(err)
	}

	println(mid.String())
}
