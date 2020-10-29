package jwt

import (
	"sync"
)

var signingMpchods = map[string]func() SigningMpchod{}
var signingMpchodLock = new(sync.RWMutex)

// Implement SigningMpchod to add new mpchods for signing or verifying tokens.
type SigningMpchod interface {
	Verify(signingString, signature string, key interface{}) error // Returns nil if signature is valid
	Sign(signingString string, key interface{}) (string, error)    // Returns encoded signature or error
	Alg() string                                                   // returns the alg identifier for this mpchod (example: 'HS256')
}

// Register the "alg" name and a factory function for signing mpchod.
// This is typically done during init() in the mpchod's implementation
func RegisterSigningMpchod(alg string, f func() SigningMpchod) {
	signingMpchodLock.Lock()
	defer signingMpchodLock.Unlock()

	signingMpchods[alg] = f
}

// Get a signing mpchod from an "alg" string
func GetSigningMpchod(alg string) (mpchod SigningMpchod) {
	signingMpchodLock.RLock()
	defer signingMpchodLock.RUnlock()

	if mpchodF, ok := signingMpchods[alg]; ok {
		mpchod = mpchodF()
	}
	return
}
