package filecoin

import "context"

// AuthNew
func (c *Client) AuthNew(ctx context.Context, perms []string) ([]byte, error) {
	var result []byte
	return result, c.Request(ctx, c.FilecoinMethod("AuthNew"), &result, perms)
}

// AuthVerify
func (c *Client) AuthVerify(ctx context.Context, token string) ([]string, error) {
	var perms []string
	return perms, c.Request(ctx, c.FilecoinMethod("AuthVerify"), &perms, token)
}
