package filecoin

import "github.com/shopspring/decimal"

// 将大数的Fil转换为小数
func ToFil(v decimal.Decimal) decimal.Decimal {
	return v.Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(18)))
}

// 将小数的Fil转换为大数
func FromFil(v decimal.Decimal) decimal.Decimal {
	return v.Mul(decimal.NewFromInt(10).Pow(decimal.NewFromInt(18))).Truncate(0)
}
