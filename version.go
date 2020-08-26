package filecoin

import "context"

// Version Get Lotus Node version
func (c *Client) Version(ctx context.Context) (*Version, error) {
	var version Version
	return &version, c.Request(ctx, c.FilecoinMethod("Version"), &version)
}
