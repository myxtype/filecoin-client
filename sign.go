// 警告：目前仅支持SigTypeSecp256k1

package filecoin

import (
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/myxtype/filecoin-client/crypto"
	"github.com/myxtype/filecoin-client/lib/sigs"
	_ "github.com/myxtype/filecoin-client/lib/sigs/secp"
	"github.com/myxtype/filecoin-client/types"
)

// 生成地址私钥
func GenerateKey(typ crypto.SigType) (*address.Address, *types.KeyInfo, error) {
	pk, err := sigs.Generate(typ)
	if err != nil {
		return nil, nil, err
	}

	name, _ := typ.Name()
	ki := &types.KeyInfo{
		Type:       name,
		PrivateKey: pk,
	}

	publicKey, err := sigs.ToPublic(typ, ki.PrivateKey)
	if err != nil {
		return nil, nil, err
	}

	var addr address.Address
	switch typ {
	case crypto.SigTypeSecp256k1:
		addr, err = address.NewSecp256k1Address(publicKey)
		if err != nil {
			return nil, nil, fmt.Errorf("converting Secp256k1 to address: %w", err)
		}
	case crypto.SigTypeBLS:
		addr, err = address.NewBLSAddress(publicKey)
		if err != nil {
			return nil, nil, fmt.Errorf("converting BLS to address: %w", err)
		}
	default:
		return nil, nil, fmt.Errorf("unknown key type")
	}

	return &addr, ki, nil
}

// 数据签名
func WalletSign(ki *types.KeyInfo, msg []byte) (*crypto.Signature, error) {
	return sigs.Sign(sigs.ActSigType(ki.Type), ki.PrivateKey, msg)
}

// 消息签名
func WalletSignMessage(ki *types.KeyInfo, msg *types.Message) (*crypto.Signature, error) {
	return nil, nil
}
