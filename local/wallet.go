package local

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/myxtype/filecoin-client/sigs"
	"github.com/myxtype/filecoin-client/types"
	"golang.org/x/xerrors"

	_ "github.com/myxtype/filecoin-client/sigs/secp"
)

// LocalWalletNew creates a new address in the wallet with the given sigType.
func LocalWalletNew(typ types.KeyType) (*types.KeyInfo, *address.Address, error) {
	ctyp := ActSigType(typ)
	if ctyp == crypto.SigTypeUnknown {
		return nil, nil, xerrors.Errorf("unknown sig type: %s", typ)
	}
	pk, err := sigs.Generate(ctyp)
	if err != nil {
		return nil, nil, err
	}

	publicKey, err := sigs.ToPublic(ctyp, pk)
	if err != nil {
		return nil, nil, err
	}

	var addr address.Address
	switch typ {
	case types.KTSecp256k1:
		addr, err = address.NewSecp256k1Address(publicKey)
		if err != nil {
			return nil, nil, xerrors.Errorf("converting Secp256k1 to address: %w", err)
		}
	case types.KTBLS:
		addr, err = address.NewBLSAddress(publicKey)
		if err != nil {
			return nil, nil, xerrors.Errorf("converting BLS to address: %w", err)
		}
	default:
		return nil, nil, xerrors.Errorf("unsupported key type: %s", typ)
	}

	return &types.KeyInfo{
		Type:       typ,
		PrivateKey: pk,
	}, &addr, nil
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
