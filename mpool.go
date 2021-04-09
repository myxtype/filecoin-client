package filecoin

import (
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"github.com/myxtype/filecoin-client/types"
)

// MpoolPush pushes a signed message to mempool.
func (c *Client) MpoolPush(ctx context.Context, sm *types.SignedMessage) (cid.Cid, error) {
	var id cid.Cid
	return id, c.Request(ctx, c.FilecoinMethod("MpoolPush"), &id, sm)
}

// MpoolGetNonce 获取指定发送账号的下一个nonce值
func (c *Client) MpoolGetNonce(ctx context.Context, address address.Address) (nonce uint64, err error) {
	return nonce, c.Request(ctx, c.FilecoinMethod("MpoolGetNonce"), &nonce, address)
}
