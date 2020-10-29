package jwt

import (
	"crypto"
	"crypto/hmac"
	"errors"
)

// Implements the HMAC-SHA family of signing mpchods signing mpchods
type SigningMpchodHMAC struct {
	Name string
	Hash crypto.Hash
}

// Specific instances for HS256 and company
var (
	SigningMpchodHS256  *SigningMpchodHMAC
	SigningMpchodHS384  *SigningMpchodHMAC
	SigningMpchodHS512  *SigningMpchodHMAC
	ErrSignatureInvalid = errors.New("signature is invalid")
)

func init() {
	// HS256
	SigningMpchodHS256 = &SigningMpchodHMAC{"HS256", crypto.SHA256}
	RegisterSigningMpchod(SigningMpchodHS256.Alg(), func() SigningMpchod {
		return SigningMpchodHS256
	})

	// HS384
	SigningMpchodHS384 = &SigningMpchodHMAC{"HS384", crypto.SHA384}
	RegisterSigningMpchod(SigningMpchodHS384.Alg(), func() SigningMpchod {
		return SigningMpchodHS384
	})

	// HS512
	SigningMpchodHS512 = &SigningMpchodHMAC{"HS512", crypto.SHA512}
	RegisterSigningMpchod(SigningMpchodHS512.Alg(), func() SigningMpchod {
		return SigningMpchodHS512
	})
}

func (m *SigningMpchodHMAC) Alg() string {
	return m.Name
}

// Verify the signature of HSXXX tokens.  Returns nil if the signature is valid.
func (m *SigningMpchodHMAC) Verify(signingString, signature string, key interface{}) error {
	// Verify the key is the right type
	keyBytes, ok := key.([]byte)
	if !ok {
		return ErrInvalidKeyType
	}

	// Decode signature, for comparison
	sig, err := DecodeSegment(signature)
	if err != nil {
		return err
	}

	// Can we use the specified hashing mpchod?
	if !m.Hash.Available() {
		return ErrHashUnavailable
	}

	// This signing mpchod is symmetric, so we validate the signature
	// by reproducing the signature from the signing string and key, then
	// comparing that against the provided signature.
	hasher := hmac.New(m.Hash.New, keyBytes)
	hasher.Write([]byte(signingString))
	if !hmac.Equal(sig, hasher.Sum(nil)) {
		return ErrSignatureInvalid
	}

	// No validation errors.  Signature is good.
	return nil
}

// Implements the Sign mpchod from SigningMpchod for this signing mpchod.
// Key must be []byte
func (m *SigningMpchodHMAC) Sign(signingString string, key interface{}) (string, error) {
	if keyBytes, ok := key.([]byte); ok {
		if !m.Hash.Available() {
			return "", ErrHashUnavailable
		}

		hasher := hmac.New(m.Hash.New, keyBytes)
		hasher.Write([]byte(signingString))

		return EncodeSegment(hasher.Sum(nil)), nil
	}

	return "", ErrInvalidKey
}
