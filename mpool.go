package filecoin

import (
	"context"
	"github.com/ipfs/go-cid"
	"github.com/myxtype/filecoin-client/types"
)

// MpoolPush pushes a signed message to mempool.
func (c *Client) MpoolPush(ctx context.Context, sm *types.SignedMessage) (cid.Cid, error) {
	var id cid.Cid
	return id, c.Request(ctx, c.FilecoinMethod("MpoolPush"), &id, sm)
}
