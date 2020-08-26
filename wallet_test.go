package filecoin

import (
	"context"
	"encoding/hex"
	"testing"
)

// 查询钱包余额
func TestClient_WalletBalance(t *testing.T) {
	c := testClient()

	b, err := c.WalletBalance(context.Background(), "t1e3soclcq34tq7wmykp7xkkmpkzjnghumm3syyay")
	if err != nil {
		t.Error(err)
	}

	t.Log(b.String())
}

func TestClient_WalletNew(t *testing.T) {
	c := testClient()

	// t1r6egk7djfy7krbw7zdswbgdhep4hge5fecwmsoi
	addr, err := c.WalletNew(context.Background(), SigTypeSecp256k1)
	if err != nil {
		t.Error(err)
	}
	t.Log(addr)

	// secp256k1 fd1d429f2e0744f5dbcc361796e1a6f5cf4b59ecca92c15c27f837401c12a3da
	ki, err := c.WalletExport(context.Background(), addr)
	if err != nil {
		t.Error(err)
	}

	t.Log(ki.Type, hex.EncodeToString(ki.PrivateKey))
}
