package local

import (
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/myxtype/filecoin-client/sigs"
	"github.com/myxtype/filecoin-client/types"
	// _ "github.com/myxtype/filecoin-client/sigs/bls"
	_ "github.com/myxtype/filecoin-client/sigs/secp"
)

// WalletNew creates a new address in the wallet with the given sigType.
func WalletNew(typ types.KeyType) (*types.KeyInfo, *address.Address, error) {
	ctyp := ActSigType(typ)
	if ctyp == crypto.SigTypeUnknown {
		return nil, nil, fmt.Errorf("unknown sig type: %s", typ)
	}

	pk, err := sigs.Generate(ctyp)
	if err != nil {
		return nil, nil, err
	}

	addr, err := WalletPrivateToAddress(ctyp, pk)
	if err != nil {
		return nil, nil, err
	}

	return &types.KeyInfo{
		Type:       typ,
		PrivateKey: pk,
	}, addr, nil
}

// WalletPrivateToAddress convert private key to public key address
func WalletPrivateToAddress(typ crypto.SigType, pk []byte) (*address.Address, error) {
	publicKey, err := sigs.ToPublic(typ, pk)
	if err != nil {
		return nil, err
	}

	var addr address.Address
	switch typ {
	case crypto.SigTypeSecp256k1:
		addr, err = address.NewSecp256k1Address(publicKey)
		if err != nil {
			return nil, fmt.Errorf("converting Secp256k1 to address: %w", err)
		}
	case crypto.SigTypeBLS:
		addr, err = address.NewBLSAddress(publicKey)
		if err != nil {
			return nil, fmt.Errorf("converting BLS to address: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported key type: %v", typ)
	}

	return &addr, nil
}

// WalletSign signs the given bytes using the KeyType and private key.
func WalletSign(typ types.KeyType, pk []byte, data []byte) (*crypto.Signature, error) {
	return sigs.Sign(ActSigType(typ), pk, data)
}

// WalletVerify verify the signed message
func WalletVerify(sig *crypto.Signature, addr address.Address, msg []byte) error {
	return sigs.Verify(sig, addr, msg)
}

// WalletSignMessage signs the given message using the given private key.
func WalletSignMessage(typ types.KeyType, pk []byte, msg *types.Message) (*types.SignedMessage, error) {
	mb, err := msg.ToStorageBlock()
	if err != nil {
		return nil, fmt.Errorf("serializing message: %w", err)
	}

	sig, err := WalletSign(typ, pk, mb.Cid().Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to sign message: %w", err)
	}

	return &types.SignedMessage{
		Message:   msg,
		Signature: sig,
	}, nil
}

func WalletVerifyMessage(sm *types.SignedMessage) error {
	return WalletVerify(sm.Signature, sm.Message.From, sm.Message.Cid().Bytes())
}

func ActSigType(typ types.KeyType) crypto.SigType {
	switch typ {
	case types.KTBLS:
		return crypto.SigTypeBLS
	case types.KTSecp256k1:
		return crypto.SigTypeSecp256k1
	default:
		return crypto.SigTypeUnknown
	}
}
