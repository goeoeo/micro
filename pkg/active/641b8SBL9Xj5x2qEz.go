
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNum641b8SBL9Xj5x2qEz = []byte("641b8SBL9Xj5x2qEz")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytes641b8SBL9Xj5x2qEz(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNum641b8SBL9Xj5x2qEz
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

