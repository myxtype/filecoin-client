package filecoin

import (
	"context"
	"github.com/ipfs/go-cid"
	"github.com/myxtype/filecoin-client/types"
)

// GasEstimateGasLimit estimates gas used by the message and returns it. It fails if message fails to execute.
func (c *Client) GasEstimateGasLimit(ctx context.Context, message *types.Message, cids []*cid.Cid) (int64, error) {
	var gasLimit int64
	return gasLimit, c.Request(ctx, c.FilecoinMethod("GasEstimateGasLimit"), &gasLimit, message, cids)
}

// GasEstimateMessageGas estimates gas values for unset message gas fields
func (c *Client) GasEstimateMessageGas(ctx context.Context, message *types.Message, spec *types.MessageSendSpec, cids []*cid.Cid) (*types.Message, error) {
	var msg *types.Message
	return msg, c.Request(ctx, c.FilecoinMethod("GasEstimateMessageGas"), &msg, message, spec, cids)
}
