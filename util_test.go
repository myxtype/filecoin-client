package filecoin

import (
	"github.com/filecoin-project/go-state-types/big"
	"github.com/shopspring/decimal"
	"testing"
)

func TestToFil(t *testing.T) {
	v := big.NewFromGo(decimal.RequireFromString("7213219999999877989641900").BigInt())
	t.Log(ToFil(v).String())
}

func TestFromFil(t *testing.T) {
	v := decimal.RequireFromString("7213219.9999998779896419")
	t.Log(FromFil(v))
}
