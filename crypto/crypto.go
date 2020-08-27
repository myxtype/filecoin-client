package crypto

import (
	"fmt"
	"math"
)

type SigType byte

const (
	SigTypeUnknown = SigType(math.MaxUint8)

	SigTypeSecp256k1 = SigType(iota) // Most
	SigTypeBLS
)

func (t SigType) Name() (string, error) {
	switch t {
	case SigTypeUnknown:
		return "unknown", nil
	case SigTypeSecp256k1:
		return "secp256k1", nil
	case SigTypeBLS:
		return "bls", nil
	default:
		return "", fmt.Errorf("invalid signature type: %d", t)
	}
}

type Signature struct {
	Type SigType `json:"Type"`
	Data []byte  `json:"Data"`
}
