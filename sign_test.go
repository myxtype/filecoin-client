package filecoin

import (
	"encoding/hex"
	"github.com/myxtype/filecoin-client/crypto"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	addr, ki, err := GenerateKey(crypto.SigTypeSecp256k1)
	if err != nil {
		t.Error(err)
	}

	// t12l7d6knanzfg5j26yyssyj7sfo3qbn4qwqgph2y
	t.Log(addr.String())
	// 6564c7d365c2605156ca42190863e5a57405a604ea0d5b526a45d0bc57c1c9c8
	t.Log(hex.EncodeToString(ki.PrivateKey))
}
