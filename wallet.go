package filecoin

import (
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/specs-actors/actors/crypto"
	"github.com/myxtype/filecoin-client/sigs"
	"github.com/myxtype/filecoin-client/types"
	"github.com/shopspring/decimal"
	"golang.org/x/xerrors"

	_ "github.com/myxtype/filecoin-client/sigs/secp" // enable secp signatures
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
func (c *Client) WalletExport(ctx context.Context, addr address.Address) (*types.KeyInfo, error) {
	var ki *types.KeyInfo
	return ki, c.Request(ctx, c.FilecoinMethod("WalletExport"), &ki, addr)
}

// WalletHas indicates whether the given address is in the wallet.
func (c *Client) WalletHas(ctx context.Context, addr address.Address) (bool, error) {
	var has bool
	return has, c.Request(ctx, c.FilecoinMethod("WalletHas"), &has, addr)
}

// WalletImport receives a KeyInfo, which includes a private key, and imports it into the wallet.
func (c *Client) WalletImport(ctx context.Context, ki *types.KeyInfo) (address.Address, error) {
	var addr address.Address
	return addr, c.Request(ctx, c.FilecoinMethod("WalletImport"), &addr, ki)
}

// WalletList lists all the addresses in the wallet.
func (c *Client) WalletList(ctx context.Context) ([]address.Address, error) {
	var addrs []address.Address
	return addrs, c.Request(ctx, c.FilecoinMethod("WalletList"), &addrs)
}

// WalletNew creates a new address in the wallet with the given sigType.
func (c *Client) WalletNew(ctx context.Context, sigType crypto.SigType) (address.Address, error) {
	var addr address.Address
	return addr, c.Request(ctx, c.FilecoinMethod("WalletNew"), &addr, sigType)
}

// WalletSetDefault marks the given address as as the default one.
func (c *Client) WalletSetDefault(ctx context.Context, addr address.Address) error {
	return c.Request(ctx, c.FilecoinMethod("WalletSetDefault"), nil, addr)
}

// WalletSign signs the given bytes using the given address.
func (c *Client) WalletSign(ctx context.Context, addr address.Address, data []byte) (*crypto.Signature, error) {
	var sig *crypto.Signature
	return sig, c.Request(ctx, c.FilecoinMethod("WalletSign"), &sig, addr, data)
}

// WalletSignMessage signs the given message using the given address.
func (c *Client) WalletSignMessage(ctx context.Context, addr address.Address, message *types.Message) (*types.SignedMessage, error) {
	var sm *types.SignedMessage
	return sm, c.Request(ctx, c.FilecoinMethod("WalletSignMessage"), &sm, addr, message)
}

// WalletVerify takes an address, a signature, and some bytes, and indicates whether the signature is valid. The address does not have to be in the wallet.
func (c *Client) WalletVerify(ctx context.Context, k string, msg []byte, sig *crypto.Signature) (bool, error) {
	var ok bool
	return ok, c.Request(ctx, c.FilecoinMethod("WalletVerify"), &ok, k, msg, sig)
}

// WalletSignMessage signs the given message with the given private key
func (c *Client) WalletSignMessageLocal(sigType crypto.SigType, privkey []byte, message *types.Message) (*types.SignedMessage, error) {
	mcid := message.Cid()

	sig, err := sigs.Sign(sigType, privkey, mcid.Bytes())
	if err != nil {
		return nil, xerrors.Errorf("failed to sign message: %w", err)
	}

	return &types.SignedMessage{
		Message:   message,
		Signature: *sig,
	}, nil
}

// WalletNewLocal generate an address locally and return the private key
func (c *Client) WalletNewLocal(typ crypto.SigType) (*address.Address, []byte, error) {
	pk, err := sigs.Generate(typ)
	if err != nil {
		return nil, nil, err
	}

	publicKey, err := sigs.ToPublic(typ, pk)
	if err != nil {
		return nil, nil, err
	}

	var addr address.Address
	switch typ {
	case crypto.SigTypeSecp256k1:
		addr, err = address.NewSecp256k1Address(publicKey)
		if err != nil {
			return nil, nil, xerrors.Errorf("converting Secp256k1 to address: %w", err)
		}
	//case crypto.SigTypeBLS:
	//	addr, err = address.NewBLSAddress(publicKey)
	//	if err != nil {
	//		return nil, nil, xerrors.Errorf("converting BLS to address: %w", err)
	//	}
	default:
		return nil, nil, xerrors.Errorf("unknown key type")
	}

	return &addr, pk, err
}
