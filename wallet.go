package filecoin

import (
	"context"
	"github.com/shopspring/decimal"
)

// WalletBalance returns the balance of the given address at the current head of the chain.
func (c *Client) WalletBalance(ctx context.Context, addr string) (decimal.Decimal, error) {
	var balance decimal.Decimal
	return balance, c.Request(ctx, c.FilecoinMethod("WalletBalance"), &balance, addr)
}

// WalletDefaultAddress returns the address marked as default in the wallet.
func (c *Client) WalletDefaultAddress(ctx context.Context) (string, error) {
	var addr string
	return addr, c.Request(ctx, c.FilecoinMethod("WalletDefaultAddress"), &addr)
}

// WalletExport returns the private key of an address in the wallet.
func (c *Client) WalletExport(ctx context.Context, addr string) (*KeyInfo, error) {
	var ki KeyInfo
	return &ki, c.Request(ctx, c.FilecoinMethod("WalletExport"), &ki, addr)
}

// WalletHas indicates whether the given address is in the wallet.
func (c *Client) WalletHas(ctx context.Context, addr string) (bool, error) {
	var has bool
	return has, c.Request(ctx, c.FilecoinMethod("WalletHas"), &has, addr)
}

// WalletImport receives a KeyInfo, which includes a private key, and imports it into the wallet.
func (c *Client) WalletImport(ctx context.Context, ki *KeyInfo) (string, error) {
	var addr string
	return addr, c.Request(ctx, c.FilecoinMethod("WalletImport"), &addr, ki)
}

// WalletList lists all the addresses in the wallet.
// todo

// WalletNew creates a new address in the wallet with the given sigType.
// 通常sigType设为1
func (c *Client) WalletNew(ctx context.Context, sigType int64) (string, error) {
	var addr string
	return addr, c.Request(ctx, c.FilecoinMethod("WalletNew"), &addr, sigType)
}

// WalletSetDefault marks the given address as as the default one.
func (c *Client) WalletSetDefault(ctx context.Context, addr string) error {
	return c.Request(ctx, c.FilecoinMethod("WalletSetDefault"), nil, addr)
}

// WalletSign signs the given bytes using the given address.
func (c *Client) WalletSign(ctx context.Context, addr string, data []byte) (*Signature, error) {
	var sig Signature
	return &sig, c.Request(ctx, c.FilecoinMethod("WalletSign"), &sig, addr, data)
}

// WalletSignMessage signs the given message using the given address.
func (c *Client) WalletSignMessage(ctx context.Context, addr string, message *Message) (*SignedMessage, error) {
	var sm SignedMessage
	return &sm, c.Request(ctx, c.FilecoinMethod("WalletSignMessage"), &sm, addr, message)
}

// todo
