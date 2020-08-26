package filecoin

import (
	"context"
	"github.com/ipfs/go-cid"
	"testing"
)

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
