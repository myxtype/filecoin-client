package local

import (
	"encoding/hex"
	"github.com/filecoin-project/go-address"
	"github.com/myxtype/filecoin-client/sigs/bls"
	"github.com/myxtype/filecoin-client/sigs/secp"
	"github.com/myxtype/filecoin-client/types"
	"testing"
)

func TestMain(t *testing.M){
	bls.Init()
	secp.Init()
	t.Run()
}

func TestWalletNew(t *testing.T) {

	address.CurrentNetwork = address.Mainnet

	ki, addr, err := WalletNew(types.KTBLS)
	if err != nil {
		t.Error(err)
	}

	t.Log(hex.EncodeToString(ki.PrivateKey))
	t.Log(addr.String())

	// baf063287d69520fe17d6065446a7b8a200799d676ba6844c8bfdca5dd3deb5c
	// f3qmpxmbmdjkqk73zacadt3dkt3j5ftyi4kim5bg3ubujes7gk3w6dnvrndos6wbec2rlpwimk76zu5lwl6sma
}

func TestWalletNew2(t *testing.T) {

	address.CurrentNetwork = address.Mainnet

	ki, addr, err := WalletNew(types.KTSecp256k1)
	if err != nil {
		t.Error(err)
	}

	t.Log(hex.EncodeToString(ki.PrivateKey))
	t.Log(addr.String())
	// 3e91d9dfb6a98f224745177e1c670fee00f7cf4f55f1576c34b6a0fae7b83c2c
	// f1yfi4yslez2hz3ori5grvv3xdo3xkibc4v6xjusy
}
