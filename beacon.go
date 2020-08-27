package filecoin

import (
	"context"
	"github.com/myxtype/filecoin-client/types"
)

// BeaconGetEntry returns the beacon entry for the given filecoin epoch. If the entry has not yet been produced, the call will block until the entry becomes available
func (c *Client) BeaconGetEntry(ctx context.Context, epoch int64) (*types.BeaconEntry, error) {
	var be *types.BeaconEntry
	return be, c.Request(ctx, c.FilecoinMethod("BeaconGetEntry"), &be, epoch)
}
