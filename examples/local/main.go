package main

import (
	"context"
	"encoding/hex"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/myxtype/filecoin-client"
	"github.com/myxtype/filecoin-client/local"
	"github.com/myxtype/filecoin-client/types"
	"github.com/shopspring/decimal"
)

func main() {
	// 设置网络类型
	address.CurrentNetwork = address.Mainnet

	// 生产新的地址
	// 新地址有转入fil才激活，不然没法用
	ki, addr, err := local.WalletNew(types.KTSecp256k1)
	if err != nil {
		panic(err)
	}

	println(hex.EncodeToString(ki.PrivateKey))
	println(addr.String())

	// 50a5e6234f5fdfc026bd889347409e11b6ef5b6034a7b0572d7e24ed1e9ba0e4
	// f1dynqskhlixt5eswpff3a72ksprqmeompv3pbesy

	to, _ := address.NewFromString("f1yfi4yslez2hz3ori5grvv3xdo3xkibc4v6xjusy")

	// 转移0.001FIL到f1yfi4yslez2hz3ori5grvv3xdo3xkibc4v6xjusy
	msg := &types.Message{
		Version:    0,
		To:         to,
		From:       *addr,
		Nonce:      0,
		Value:      filecoin.FromFil(decimal.NewFromFloat(0.001)),
		GasLimit:   0,
		GasFeeCap:  abi.NewTokenAmount(10000),
		GasPremium: abi.NewTokenAmount(10000),
		Method:     0,
		Params:     nil,
	}

	client := filecoin.New("https://1lB5G4SmGdSTikOo7l6vYlsktdd:b58884915362a99b4fc18c2bf8af8358@filecoin.infura.io")

	// 最大手续费0.0001 FIL
	//maxFee := filecoin.FromFil(decimal.NewFromFloat(0.0001))

	// 估算GasLimit
	//msg, err = client.GasEstimateMessageGas(context.Background(), msg, &types.MessageSendSpec{MaxFee: maxFee}, nil)
	//if err != nil {
	//	panic(err)
	//}

	// 离线签名
	s, err := local.WalletSignMessage(types.KTSecp256k1, ki.PrivateKey, msg)
	if err != nil {
		panic(err)
	}

	println(hex.EncodeToString(s.Signature.Data))
	// 47bcbb167fd9040bd02dba02789bc7bc0463c290db1be9b07065c12a64fb84dc546bef7aedfba789d0d7ce2c4532f8fa0d2dd998985ad3ec1a8b064c26e4625a01

	// 验证签名
	if err := local.WalletVerifyMessage(s); err != nil {
		panic(err)
	}

	mid, err := client.MpoolPush(context.Background(), s)
	if err != nil {
		panic(err)
	}

	println(mid.String())
}
