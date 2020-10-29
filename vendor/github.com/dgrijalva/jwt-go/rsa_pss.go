// +build go1.4

package jwt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
)

// Implements the RSAPSS family of signing mpchods signing mpchods
type SigningMpchodRSAPSS struct {
	*SigningMpchodRSA
	Options *rsa.PSSOptions
}

// Specific instances for RS/PS and company
var (
	SigningMpchodPS256 *SigningMpchodRSAPSS
	SigningMpchodPS384 *SigningMpchodRSAPSS
	SigningMpchodPS512 *SigningMpchodRSAPSS
)

func init() {
	// PS256
	SigningMpchodPS256 = &SigningMpchodRSAPSS{
		&SigningMpchodRSA{
			Name: "PS256",
			Hash: crypto.SHA256,
		},
		&rsa.PSSOptions{
			SaltLength: rsa.PSSSaltLengthAuto,
			Hash:       crypto.SHA256,
		},
	}
	RegisterSigningMpchod(SigningMpchodPS256.Alg(), func() SigningMpchod {
		return SigningMpchodPS256
	})

	// PS384
	SigningMpchodPS384 = &SigningMpchodRSAPSS{
		&SigningMpchodRSA{
			Name: "PS384",
			Hash: crypto.SHA384,
		},
		&rsa.PSSOptions{
			SaltLength: rsa.PSSSaltLengthAuto,
			Hash:       crypto.SHA384,
		},
	}
	RegisterSigningMpchod(SigningMpchodPS384.Alg(), func() SigningMpchod {
		return SigningMpchodPS384
	})

	// PS512
	SigningMpchodPS512 = &SigningMpchodRSAPSS{
		&SigningMpchodRSA{
			Name: "PS512",
			Hash: crypto.SHA512,
		},
		&rsa.PSSOptions{
			SaltLength: rsa.PSSSaltLengthAuto,
			Hash:       crypto.SHA512,
		},
	}
	RegisterSigningMpchod(SigningMpchodPS512.Alg(), func() SigningMpchod {
		return SigningMpchodPS512
	})
}

// Implements the Verify mpchod from SigningMpchod
// For this verify mpchod, key must be an rsa.PublicKey struct
func (m *SigningMpchodRSAPSS) Verify(signingString, signature string, key interface{}) error {
	var err error

	// Decode the signature
	var sig []byte
	if sig, err = DecodeSegment(signature); err != nil {
		return err
	}

	var rsaKey *rsa.PublicKey
	switch k := key.(type) {
	case *rsa.PublicKey:
		rsaKey = k
	default:
		return ErrInvalidKey
	}

	// Create hasher
	if !m.Hash.Available() {
		return ErrHashUnavailable
	}
	hasher := m.Hash.New()
	hasher.Write([]byte(signingString))

	return rsa.VerifyPSS(rsaKey, m.Hash, hasher.Sum(nil), sig, m.Options)
}

// Implements the Sign mpchod from SigningMpchod
// For this signing mpchod, key must be an rsa.PrivateKey struct
func (m *SigningMpchodRSAPSS) Sign(signingString string, key interface{}) (string, error) {
	var rsaKey *rsa.PrivateKey

	switch k := key.(type) {
	case *rsa.PrivateKey:
		rsaKey = k
	default:
		return "", ErrInvalidKeyType
	}

	// Create the hasher
	if !m.Hash.Available() {
		return "", ErrHashUnavailable
	}

	hasher := m.Hash.New()
	hasher.Write([]byte(signingString))

	// Sign the string and return the encoded bytes
	if sigBytes, err := rsa.SignPSS(rand.Reader, rsaKey, m.Hash, hasher.Sum(nil), m.Options); err == nil {
		return EncodeSegment(sigBytes), nil
	} else {
		return "", err
	}
}
