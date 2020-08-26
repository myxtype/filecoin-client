package filecoin

import (
	"context"
	"encoding/hex"
	"github.com/ipfs/go-cid"
	"github.com/shopspring/decimal"
	"testing"
)

// The Lotus Node
// The default token is in ~/.lotus/token
func testClient() *Client {
	return NewClient("http://127.0.0.1:1234/rpc/v0", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.cF__3r_0IR9KwZ2nLkqcBW8vuPePruZieJAVvTAoUA4")
}

// 测试RpcClient
func TestClient_Request(t *testing.T) {
	c := NewClient("https://eth-mainnet.token.im", "")
	var blockNumber string
	if err := c.Request(context.Background(), "eth_blockNumber", &blockNumber); err != nil {
		t.Error(err)
	}

	t.Log(blockNumber)

	var tr struct {
		BlockHash   string `json:"blockHash"`
		BlockNumber string `json:"blockNumber"`
	}
	if err := c.Request(context.Background(), "eth_getTransactionReceipt", &tr, "0xbb3a336e3f823ec18197f1e13ee875700f08f03e2cab75f0d0b118dabb44cba0"); err != nil {
		t.Error(err)
	}

	t.Log(tr.BlockHash)
	t.Log(tr.BlockNumber)
}

// 查询钱包余额
func TestClient_WalletBalance(t *testing.T) {
	c := testClient()

	b, err := c.WalletBalance(context.Background(), "t1e3soclcq34tq7wmykp7xkkmpkzjnghumm3syyay")
	if err != nil {
		t.Error(err)
	}

	t.Log(b.String())
}

// 根据消息Cid获取消息
func TestClient_ChainGetMessage(t *testing.T) {
	c := testClient()

	id, err := cid.Parse("bafy2bzacec2iz32npq2ytfj4d4eeppidpbqjfd3glrrset7cham7sfpzinq24")
	if err != nil {
		t.Error(err)
	}

	msg, err := c.ChainGetMessage(context.Background(), id)
	if err != nil {
		t.Error(err)
	}

	t.Log(msg)
}

// 查询消息/交易执行状态
func TestClient_StateGetReceipt(t *testing.T) {
	c := testClient()

	id, err := cid.Parse("bafy2bzacea64rujzydbpzglzbe2gs74v2p2tcxgvyuzupb6fzcydrcssqg4fa")
	if err != nil {
		t.Error(err)
	}

	mr, err := c.StateGetReceipt(context.Background(), id, nil)
	if err != nil {
		t.Error(err)
	}

	t.Log(mr)
}

// 获取当前头部高度
func TestClient_ChainHead(t *testing.T) {
	c := testClient()

	ts, err := c.ChainHead(context.Background())
	if err != nil {
		t.Error(err)
	}

	t.Log(ts.Height)

	for _, n := range ts.Cids {
		bm, err := c.ChainGetBlockMessages(context.Background(), n)
		if err != nil {
			t.Error(err)
		}
		for index, msg := range bm.BlsMessages {
			t.Log(bm.Cids[index], msg)
		}
	}
}

// 根据高度遍历区块所有交易
func TestClient_ChainGetTipSetByHeight(t *testing.T) {
	c := testClient()

	ts, err := c.ChainGetTipSetByHeight(context.Background(), 1201, nil)
	if err != nil {
		t.Error(err)
	}
	for _, n := range ts.Cids {
		bm, err := c.ChainGetBlockMessages(context.Background(), n)
		if err != nil {
			t.Error(err)
		}
		for index, msg := range bm.BlsMessages {
			t.Log(bm.Cids[index], msg)
		}
	}
}

func TestClient_WalletNew(t *testing.T) {
	c := testClient()

	// t1r6egk7djfy7krbw7zdswbgdhep4hge5fecwmsoi
	addr, err := c.WalletNew(context.Background(), 1)
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

// 发送FileCoin
func TestClient_MpoolPush(t *testing.T) {
	c := testClient()

	msg := &Message{
		Version:    0,
		To:         "t1r6egk7djfy7krbw7zdswbgdhep4hge5fecwmsoi",
		From:       "t1e3soclcq34tq7wmykp7xkkmpkzjnghumm3syyay",
		Nonce:      0,
		Value:      decimal.RequireFromString("10000000000000000000"),
		GasLimit:   0,
		GasFeeCap:  decimal.Zero,
		GasPremium: decimal.Zero,
		Method:     0,
		Params:     nil,
	}

	msg, err := c.GasEstimateMessageGas(context.Background(), msg, &MessageSendSpec{MaxFee: "1000000000000"}, nil)
	if err != nil {
		t.Error(err)
	}

	actor, err := c.StateGetActor(context.Background(), msg.From, nil)
	if err != nil {
		t.Error(err)
	}

	msg.Nonce = actor.Nonce
	t.Log(msg)

	sm, err := c.WalletSignMessage(context.Background(), msg.From, msg)
	if err != nil {
		t.Error(err)
	}

	id, err := c.MpoolPush(context.Background(), sm)
	if err != nil {
		t.Error(err)
	}

	t.Log(id.Version())
	t.Log(id)
}
