package filecoin

import (
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/shopspring/decimal"
)

// WalletBalance returns the balance of the given address at the current head of the chain.
func (c *Client) WalletBalance(ctx context.Context, addr address.Address) (decimal.Decimal, error) {
	var balance decimal.Decimal
	return balance, c.Request(ctx, c.FilecoinMethod("WalletBalance"), &balance, addr)
}

// WalletDefaultAddress returns the address marked as default in the wallet.
func (c *Client) WalletDefaultAddress(ctx context.Context) (address.Address, error) {
	var addr address.Address
	return addr, c.Request(ctx, c.FilecoinMethod("WalletDefaultAddress"), &addr)
}

// WalletDelete deletes an address from the wallet.
func (c *Client) WalletDelete(ctx context.Context, addr address.Address) error {
	return c.Request(ctx, c.FilecoinMethod("WalletDelete"), nil, addr)
}

// WalletExport returns the private key of an address in the wallet.
func (c *Client) WalletExport(ctx context.Context, addr address.Address) (*KeyInfo, error) {
	var ki *KeyInfo
	return ki, c.Request(ctx, c.FilecoinMethod("WalletExport"), &ki, addr)
}

// WalletHas indicates whether the given address is in the wallet.
func (c *Client) WalletHas(ctx context.Context, addr address.Address) (bool, error) {
	var has bool
	return has, c.Request(ctx, c.FilecoinMethod("WalletHas"), &has, addr)
}

// WalletImport receives a KeyInfo, which includes a private key, and imports it into the wallet.
func (c *Client) WalletImport(ctx context.Context, ki *KeyInfo) (address.Address, error) {
	var addr address.Address
	return addr, c.Request(ctx, c.FilecoinMethod("WalletImport"), &addr, ki)
}

// WalletList lists all the addresses in the wallet.
func (c *Client) WalletList(ctx context.Context) ([]address.Address, error) {
	var addrs []address.Address
	return addrs, c.Request(ctx, c.FilecoinMethod("WalletList"), &addrs)
}

// WalletNew creates a new address in the wallet with the given sigType.
func (c *Client) WalletNew(ctx context.Context, sigType SigType) (address.Address, error) {
	var addr address.Address
	return addr, c.Request(ctx, c.FilecoinMethod("WalletNew"), &addr, sigType)
}

// WalletSetDefault marks the given address as as the default one.
func (c *Client) WalletSetDefault(ctx context.Context, addr address.Address) error {
	return c.Request(ctx, c.FilecoinMethod("WalletSetDefault"), nil, addr)
}

// WalletSign signs the given bytes using the given address.
func (c *Client) WalletSign(ctx context.Context, addr address.Address, data []byte) (*Signature, error) {
	var sig *Signature
	return sig, c.Request(ctx, c.FilecoinMethod("WalletSign"), &sig, addr, data)
}

// WalletSignMessage signs the given message using the given address.
func (c *Client) WalletSignMessage(ctx context.Context, addr address.Address, message *Message) (*SignedMessage, error) {
	var sm *SignedMessage
	return sm, c.Request(ctx, c.FilecoinMethod("WalletSignMessage"), &sm, addr, message)
}

// WalletVerify takes an address, a signature, and some bytes, and indicates whether the signature is valid. The address does not have to be in the wallet.
func (c *Client) WalletVerify(ctx context.Context, k string, msg []byte, sig *Signature) (bool, error) {
	var ok bool
	return ok, c.Request(ctx, c.FilecoinMethod("WalletVerify"), &ok, k, msg, sig)
}
