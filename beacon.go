package filecoin

import "context"

// BeaconGetEntry returns the beacon entry for the given filecoin epoch. If the entry has not yet been produced, the call will block until the entry becomes available
func (c *Client) BeaconGetEntry(ctx context.Context, epoch int64) (*BeaconEntry, error) {
	var be *BeaconEntry
	return be, c.Request(ctx, c.FilecoinMethod("BeaconGetEntry"), &be, epoch)
}
