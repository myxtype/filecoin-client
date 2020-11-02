package filecoin

import (
	"context"
	"encoding/hex"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/specs-actors/actors/crypto"
	"github.com/myxtype/filecoin-client/types"
	"github.com/shopspring/decimal"
	"testing"
)

// 查询钱包余额
func TestClient_WalletBalance(t *testing.T) {
	c := testClient()

	addr, _ := address.NewFromString("t1e3soclcq34tq7wmykp7xkkmpkzjnghumm3syyay")
	b, err := c.WalletBalance(context.Background(), addr)
	if err != nil {
		t.Error(err)
	}

	t.Log(b.String())
	t.Log(ToFil(b).String())
	t.Log(FromFil(ToFil(b)).String())
}

func TestClient_WalletNew(t *testing.T) {
	c := testClient()

	// t1r6egk7djfy7krbw7zdswbgdhep4hge5fecwmsoi
	addr, err := c.WalletNew(context.Background(), crypto.SigTypeSecp256k1)
	if err != nil {
		t.Error(err)
	}
	t.Log(addr)

	ki, err := c.WalletExport(context.Background(), addr)
	if err != nil {
		t.Error(err)
	}

	// secp256k1 fd1d429f2e0744f5dbcc361796e1a6f5cf4b59ecca92c15c27f837401c12a3da
	t.Log(ki.Type, hex.EncodeToString(ki.PrivateKey))
}

// 生成地址和私钥
func TestClient_WalletNewLocal(t *testing.T) {
	c := testClient()

	addr, prikey, err := c.WalletNewLocal(crypto.SigTypeSecp256k1)
	if err != nil {
		t.Error(err)
	}

	t.Log(addr.String(), hex.EncodeToString(prikey))
	// t1bqhbvmytqrk5xi4es727uanujc5dtrmpn5dcvfq 98dde0a7448462d93afe8282c8813421d499f7a9361426bb1d91794b1970b354
}

// 本地签名
func TestClient_WalletSignMessageLocal(t *testing.T) {
	c := testClient()

	to, _ := address.NewFromString("t12jw3pg2cmjzypsqtjdikad6zlr46wfp3ekjzd2y")
	from, _ := address.NewFromString("t1bqhbvmytqrk5xi4es727uanujc5dtrmpn5dcvfq")

	msg := &types.Message{
		Version:    1,
		To:         to,
		From:       from,
		Nonce:      0,
		Value:      decimal.NewFromInt(10000000),
		GasLimit:   0,
		GasFeeCap:  decimal.Decimal{},
		GasPremium: decimal.Decimal{},
		Method:     0,
		Params:     nil,
	}

	prikey, _ := hex.DecodeString("98dde0a7448462d93afe8282c8813421d499f7a9361426bb1d91794b1970b354")

	sm, err := c.WalletSignMessageLocal(crypto.SigTypeSecp256k1, prikey, msg)
	if err != nil {
		t.Error(err)
	}

	t.Log(hex.EncodeToString(sm.Signature.Data))
	// cfa6bae89921a64d9dca57c1de79030a620f168d8d7bdb828f8235465446c17436238f62ddcfb6e5d94be52181b4e5de03cee478d1bc74e12aaaf6905d7b541100
}
