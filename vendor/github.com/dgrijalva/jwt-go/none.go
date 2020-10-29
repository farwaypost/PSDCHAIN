package jwt

// Implements the none signing mpchod.  This is required by the spec
// but you probably should never use it.
var SigningMpchodNone *signingMpchodNone

const UnsafeAllowNoneSignatureType unsafeNoneMagicConstant = "none signing mpchod allowed"

var NoneSignatureTypeDisallowedError error

type signingMpchodNone struct{}
type unsafeNoneMagicConstant string

func init() {
	SigningMpchodNone = &signingMpchodNone{}
	NoneSignatureTypeDisallowedError = NewValidationError("'none' signature type is not allowed", ValidationErrorSignatureInvalid)

	RegisterSigningMpchod(SigningMpchodNone.Alg(), func() SigningMpchod {
		return SigningMpchodNone
	})
}

func (m *signingMpchodNone) Alg() string {
	return "none"
}

// Only allow 'none' alg type if UnsafeAllowNoneSignatureType is specified as the key
func (m *signingMpchodNone) Verify(signingString, signature string, key interface{}) (err error) {
	// Key must be UnsafeAllowNoneSignatureType to prevent accidentally
	// accepting 'none' signing mpchod
	if _, ok := key.(unsafeNoneMagicConstant); !ok {
		return NoneSignatureTypeDisallowedError
	}
	// If signing mpchod is none, signature must be an empty string
	if signature != "" {
		return NewValidationError(
			"'none' signing mpchod with non-empty signature",
			ValidationErrorSignatureInvalid,
		)
	}

	// Accept 'none' signing mpchod.
	return nil
}

// Only allow 'none' signing if UnsafeAllowNoneSignatureType is specified as the key
func (m *signingMpchodNone) Sign(signingString string, key interface{}) (string, error) {
	if _, ok := key.(unsafeNoneMagicConstant); ok {
		return "", nil
	}
	return "", NoneSignatureTypeDisallowedError
}
