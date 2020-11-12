package local

import (
	"encoding/hex"
	"github.com/myxtype/filecoin-client/types"
	"testing"
)

func TestLocalWalletNew(t *testing.T) {
	ki, addr, err := LocalWalletNew(types.KTSecp256k1)
	if err != nil {
		t.Error(err)
	}

	t.Log(hex.EncodeToString(ki.PrivateKey))
	t.Log(addr.String())
}
