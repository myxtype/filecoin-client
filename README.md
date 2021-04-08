# filecoin-client

[![goproxy.cn](https://goproxy.cn/stats/github.com/myxtype/filecoin-client/badges/download-count.svg)](https://goproxy.cn)

需要自行部署Lotus Node节点：https://lotu.sh/en+getting-started

此库仅添加部分方法，但已经满足钱包充值提现逻辑了，如果需要其他方法，请Fork后自行添加。

充值流程：获取头部高度，从本地高度遍历到头部高度，再根据高度获取区块CID，根据区块CID获取区块的所有消息，判断消息的类型是否0(0为发送Fil)，和接收地址是否是本地的地址。

说明请查询client_test文件。

# 安装

`go get github.com/myxtype/filecoin-client`

# 使用

```go
package main

import (
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/myxtype/filecoin-client"
)

func main() {
	client := filecoin.NewClient("http://127.0.0.1:1234/rpc/v0", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.cF__3r_0IR9KwZ2nLkqcBW8vuPePruZieJAVvTAoUA4")

	addr, _ := address.NewFromString("t1e3soclcq34tq7wmykp7xkkmpkzjnghumm3syyay")
	bal, err := client.WalletBalance(context.Background(), addr)
	if err != nil {
		panic(err)
	}

	println(bal.String())
}
```

# 离线签名版

公共节点申请：https://infura.io

具体实现逻辑在官方库中：https://github.com/filecoin-project/go-filecoin

```go
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

```

> 暂时不支持`bls`类型，仅支持`secp256k1`，所以离线签名所有类型请选择`types.KTSecp256k1`

> 私钥存储在数据库请进行加密存储，建议进行AES加密，将AES密钥放在代码中，别放配置文件，编译到执行文件中。

# Util

util.go 中提供FIL小数与大数之间的转换，但没有严格测试，请小心使用。

# 充值流程

参见：/examples/recharge

# Lotus文档

https://lotu.sh/en+api-methods
