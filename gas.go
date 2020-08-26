package filecoin

import (
	"context"
	"github.com/ipfs/go-cid"
)

// GasEstimateGasLimit estimates gas used by the message and returns it. It fails if message fails to execute.
func (c *Client) GasEstimateGasLimit(ctx context.Context, message *Message, cids []*cid.Cid) (int64, error) {
	var gasLimit int64
	return gasLimit, c.Request(ctx, c.FilecoinMethod("GasEstimateGasLimit"), &gasLimit, message, cids)
}

// GasEstimateMessageGas estimates gas values for unset message gas fields
func (c *Client) GasEstimateMessageGas(ctx context.Context, message *Message, spec *MessageSendSpec, cids []*cid.Cid) (*Message, error) {
	var msg Message
	return &msg, c.Request(ctx, c.FilecoinMethod("GasEstimateMessageGas"), &msg, message, spec, cids)
}
