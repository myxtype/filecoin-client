package filecoin

import (
	"context"
	"encoding/hex"
	"github.com/filecoin-project/go-address"
	"github.com/myxtype/filecoin-client/types"
	"testing"
)

// 查询钱包余额
func TestClient_WalletBalance(t *testing.T) {
	c := testClient()

	addr, _ := address.NewFromString("f1ntod647g54mv7pqbkniqnyov6k7thr2uxdec42i")
	b, err := c.WalletBalance(context.Background(), addr)
	if err != nil {
		t.Error(err)
	}

	t.Log(ToFil(b))
}

func TestClient_WalletNew(t *testing.T) {
	c := testClient()

	// t1r6egk7djfy7krbw7zdswbgdhep4hge5fecwmsoi
	addr, err := c.WalletNew(context.Background(), types.KTSecp256k1)
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
