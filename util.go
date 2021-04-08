package filecoin

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/shopspring/decimal"
)

// 将大数的Fil转换为小数
func ToFil(v abi.TokenAmount) decimal.Decimal {
	d := decimal.NewFromBigInt(v.Int, 0)
	return d.DivRound(decimal.NewFromInt(10).Pow(decimal.NewFromInt(18)), 18)
}

// 将小数的Fil转换为大数
func FromFil(v decimal.Decimal) abi.TokenAmount {
	return big.NewFromGo(v.Mul(decimal.NewFromInt(10).Pow(decimal.NewFromInt(18))).BigInt())
}
