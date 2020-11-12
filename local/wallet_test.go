package local

import (
	"encoding/hex"
	"github.com/filecoin-project/go-address"
	"github.com/myxtype/filecoin-client/types"
	"testing"
)

func TestLocalWalletNew(t *testing.T) {
	// 指定节点网络
	address.CurrentNetwork = address.Mainnet
	// 生成地址对
	ki, addr, err := LocalWalletNew(types.KTBLS)
	if err != nil {
		t.Error(err)
	}

	t.Log(hex.EncodeToString(ki.PrivateKey))
	t.Log(addr.String())

	// baf063287d69520fe17d6065446a7b8a200799d676ba6844c8bfdca5dd3deb5c
	// f3qmpxmbmdjkqk73zacadt3dkt3j5ftyi4kim5bg3ubujes7gk3w6dnvrndos6wbec2rlpwimk76zu5lwl6sma
}

func TestLocalWalletNew2(t *testing.T) {
	// 指定节点网络
	address.CurrentNetwork = address.Mainnet
	// 生成地址对
	ki, addr, err := LocalWalletNew(types.KTSecp256k1)
	if err != nil {
		t.Error(err)
	}

	t.Log(hex.EncodeToString(ki.PrivateKey))
	t.Log(addr.String())
}
