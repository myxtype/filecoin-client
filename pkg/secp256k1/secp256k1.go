package secp256k1

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/myxtype/filecoin-client/pkg/btcec"
	"io"
)

// PrivateKeyBytes is the size of a serialized private key.
const PrivateKeyBytes = 32

// GenerateKey creates a new key using secure randomness from crypto.rand.
func GenerateKey() ([]byte, error) {
	return GenerateKeyFromSeed(rand.Reader)
}

// GenerateKeyFromSeed generates a new key from the given reader.
func GenerateKeyFromSeed(seed io.Reader) ([]byte, error) {
	key, err := ecdsa.GenerateKey(btcec.S256(), seed)
	if err != nil {
		return nil, err
	}

	privkey := make([]byte, PrivateKeyBytes)
	blob := key.D.Bytes()

	// the length is guaranteed to be fixed, given the serialization rules for secp2561k curve points.
	copy(privkey[PrivateKeyBytes-len(blob):], blob)

	return privkey, nil
}

// PublicKey returns the public key for this private key.
func PublicKey(sk []byte) []byte {
	x, y := btcec.S256().ScalarBaseMult(sk)
	return elliptic.Marshal(btcec.S256(), x, y)
}

// Sign signs the given message, which must be 32 bytes long.
func Sign(sk, msg []byte) ([]byte, error) {
	p, _ := btcec.PrivKeyFromBytes(btcec.S256(), sk)

	sig, err := btcec.SignCompact(btcec.S256(), p, msg, false)
	if err != nil {
		return nil, err
	}

	v := sig[0] - 27
	copy(sig, sig[1:])
	sig[64] = v

	return sig, nil
}

// EcRecover recovers the public key from a message, signature pair.
func EcRecover(msg, signature []byte) ([]byte, error) {
	var sig = make([]byte, 65)
	copy(sig, signature)

	v := sig[64] + 27
	copy(sig[1:], sig[:64])
	sig[0] = v

	pk, _, err := btcec.RecoverCompact(btcec.S256(), sig, msg)
	if err != nil {
		return nil, err
	}

	return pk.SerializeUncompressed(), nil
}
