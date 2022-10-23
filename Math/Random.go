package Math

import (
	crand "crypto/rand"
	"log"
	"math/big"
	"math/rand"
)

// GetRand returns a cryptographically secure random number source
func GetRand() *rand.Rand {
	// get random seed from crypto/rand
	cnum, err := crand.Int(crand.Reader, big.NewInt(1<<63-1))
	if err != nil {
		log.Panic(err)
	}
	return rand.New(rand.NewSource(cnum.Int64()))
}
