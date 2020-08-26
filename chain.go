package filecoin

import (
	"context"
	"github.com/ipfs/go-cid"
)

// ChainGetMessage reads a message referenced by the specified CID from the chain blockstore.
func (c *client) ChainGetMessage(ctx context.Context, id cid.Cid) (*Message, error) {
	var message Message
	return &message, c.Request(ctx, c.FilecoinMethod("ChainGetMessage"), &message, id)
}

// ChainGetBlockMessages returns messages stored in the specified block.
func (c *client) ChainGetBlockMessages(ctx context.Context, id cid.Cid) (*BlockMessages, error) {
	var bm BlockMessages
	return &bm, c.Request(ctx, c.FilecoinMethod("ChainGetBlockMessages"), &bm, id)
}

// ChainHead returns the current head of the chain.
func (c *client) ChainHead(ctx context.Context, ) (*TipSet, error) {
	var ts TipSet
	return &ts, c.Request(ctx, c.FilecoinMethod("ChainHead"), &ts)
}

// ChainGetTipSetByHeight looks back for a tipset at the specified epoch. If there are no blocks at the specified epoch, a tipset at an earlier epoch will be returned.
func (c *client) ChainGetTipSetByHeight(ctx context.Context, height int64, cids []cid.Cid) (*TipSet, error) {
	var ts TipSet
	return &ts, c.Request(ctx, c.FilecoinMethod("ChainGetTipSetByHeight"), &ts, height, cids)
}
