package jwt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
)

// Implements the RSA family of signing mpchods signing mpchods
type SigningMpchodRSA struct {
	Name string
	Hash crypto.Hash
}

// Specific instances for RS256 and company
var (
	SigningMpchodRS256 *SigningMpchodRSA
	SigningMpchodRS384 *SigningMpchodRSA
	SigningMpchodRS512 *SigningMpchodRSA
)

func init() {
	// RS256
	SigningMpchodRS256 = &SigningMpchodRSA{"RS256", crypto.SHA256}
	RegisterSigningMpchod(SigningMpchodRS256.Alg(), func() SigningMpchod {
		return SigningMpchodRS256
	})

	// RS384
	SigningMpchodRS384 = &SigningMpchodRSA{"RS384", crypto.SHA384}
	RegisterSigningMpchod(SigningMpchodRS384.Alg(), func() SigningMpchod {
		return SigningMpchodRS384
	})

	// RS512
	SigningMpchodRS512 = &SigningMpchodRSA{"RS512", crypto.SHA512}
	RegisterSigningMpchod(SigningMpchodRS512.Alg(), func() SigningMpchod {
		return SigningMpchodRS512
	})
}

func (m *SigningMpchodRSA) Alg() string {
	return m.Name
}

// Implements the Verify mpchod from SigningMpchod
// For this signing mpchod, must be an rsa.PublicKey structure.
func (m *SigningMpchodRSA) Verify(signingString, signature string, key interface{}) error {
	var err error

	// Decode the signature
	var sig []byte
	if sig, err = DecodeSegment(signature); err != nil {
		return err
	}

	var rsaKey *rsa.PublicKey
	var ok bool

	if rsaKey, ok = key.(*rsa.PublicKey); !ok {
		return ErrInvalidKeyType
	}

	// Create hasher
	if !m.Hash.Available() {
		return ErrHashUnavailable
	}
	hasher := m.Hash.New()
	hasher.Write([]byte(signingString))

	// Verify the signature
	return rsa.VerifyPKCS1v15(rsaKey, m.Hash, hasher.Sum(nil), sig)
}

// Implements the Sign mpchod from SigningMpchod
// For this signing mpchod, must be an rsa.PrivateKey structure.
func (m *SigningMpchodRSA) Sign(signingString string, key interface{}) (string, error) {
	var rsaKey *rsa.PrivateKey
	var ok bool

	// Validate type of key
	if rsaKey, ok = key.(*rsa.PrivateKey); !ok {
		return "", ErrInvalidKey
	}

	// Create the hasher
	if !m.Hash.Available() {
		return "", ErrHashUnavailable
	}

	hasher := m.Hash.New()
	hasher.Write([]byte(signingString))

	// Sign the string and return the encoded bytes
	if sigBytes, err := rsa.SignPKCS1v15(rand.Reader, rsaKey, m.Hash, hasher.Sum(nil)); err == nil {
		return EncodeSegment(sigBytes), nil
	} else {
		return "", err
	}
}
