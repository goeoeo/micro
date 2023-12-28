
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumpqg7YGdoyGe76o = []byte("pqg7YGdoyGe76o")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytespqg7YGdoyGe76o(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumpqg7YGdoyGe76o
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

