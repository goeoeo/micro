
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumA9yPtxtoVYojACL1SSUwlGyDTnN0qJ = []byte("A9yPtxtoVYojACL1SSUwlGyDTnN0qJ")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesA9yPtxtoVYojACL1SSUwlGyDTnN0qJ(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumA9yPtxtoVYojACL1SSUwlGyDTnN0qJ
	}
	var bytes = make([]byte, n)
	var randBy bool
	if num, err := rand.Read(bytes); num != n || err != nil {
		r.Seed(time.Now().UnixNano())
		randBy = true
	}
	for i, b := range bytes {
		if randBy {
			bytes[i] = alphabets[r.Intn(len(alphabets))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return bytes
}

