package filecoin

import (
	"context"
	"github.com/ipfs/go-cid"
)

// StateGetActor returns the indicated actor's nonce and balance.
func (c *client) StateGetActor(ctx context.Context, addr string, cids []*cid.Cid) (*Actor, error) {
	var actor Actor
	return &actor, c.Request(ctx, c.FilecoinMethod("StateGetActor"), &actor, addr, cids)
}

// StateGetReceipt returns the message receipt for the given message
func (c *client) StateGetReceipt(ctx context.Context, id cid.Cid, cids []*cid.Cid) (*MessageReceipt, error) {
	var mr MessageReceipt
	return &mr, c.Request(ctx, c.FilecoinMethod("StateGetReceipt"), &mr, id, cids)
}
