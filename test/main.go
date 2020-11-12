package main

import (
	"encoding/hex"
	"github.com/myxtype/filecoin-client/local"
	"github.com/myxtype/filecoin-client/types"
)

func main() {
	ki, addr, err := local.LocalWalletNew(types.KTSecp256k1)
	if err != nil {
		panic(err)
	}

	println(hex.EncodeToString(ki.PrivateKey))
	println(addr.String())
}
