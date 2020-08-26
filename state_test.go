package filecoin

import (
	"context"
	"github.com/ipfs/go-cid"
	"testing"
)

// 查询消息/交易执行状态
func TestClient_StateGetReceipt(t *testing.T) {
	c := testClient()

	id, err := cid.Parse("bafy2bzacebrx3sb5do2b7cqgsnys2lkxtdq3pvjtgmdt2wclwmrtjeraa7x3q")
	if err != nil {
		t.Error(err)
	}

	mr, err := c.StateGetReceipt(context.Background(), id, nil)
	if err != nil {
		t.Error(err)
	}

	t.Log(mr)
}

// 查询消息状态
// Receipt 为空表示未执行
func TestClient_StateSearchMsg(t *testing.T) {
	c := testClient()

	id, err := cid.Parse("bafy2bzacebrx3sb5do2b7cqgsnys2lkxtdq3pvjtgmdt2wclwmrtjeraa7x3q")
	if err != nil {
		t.Error(err)
	}

	msg, err := c.StateSearchMsg(context.Background(), id)
	if err != nil {
		t.Error(err)
	}

	if msg == nil {
		t.Log("nil")
	} else {
		t.Log(msg)
	}
}
