package filecoin

import (
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"github.com/myxtype/filecoin-client/types"
)

// StateGetActor returns the indicated actor's nonce and balance.
func (c *Client) StateGetActor(ctx context.Context, addr address.Address, cids []*cid.Cid) (*types.Actor, error) {
	var actor *types.Actor
	return actor, c.Request(ctx, c.FilecoinMethod("StateGetActor"), &actor, addr, cids)
}

// StateGetReceipt returns the message receipt for the given message
func (c *Client) StateGetReceipt(ctx context.Context, id cid.Cid, cids []*cid.Cid) (*types.MessageReceipt, error) {
	var mr *types.MessageReceipt
	return mr, c.Request(ctx, c.FilecoinMethod("StateGetReceipt"), &mr, id, cids)
}

// StateReplay returns the result of executing the indicated message, assuming it was executed in the indicated tipset.
func (c *Client) StateReplay(ctx context.Context, tsk types.TipSetKey, mc cid.Cid) (*types.InvocResult, error) {
	var result *types.InvocResult
	return result, c.Request(ctx, c.FilecoinMethod("StateReplay"), &result, tsk, mc)
}

// StateSearchMsg searches for a message in the chain, and returns its receipt and the tipset where it was executed
func (c *Client) StateSearchMsg(ctx context.Context, msg cid.Cid) (*types.MsgLookup, error) {
	var msgl *types.MsgLookup
	return msgl, c.Request(ctx, c.FilecoinMethod("StateSearchMsg"), &msgl, msg)
}
