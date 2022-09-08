
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumxFQvN2P6UDmj430f62Olv44VZ = []byte("xFQvN2P6UDmj430f62Olv44VZ")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesxFQvN2P6UDmj430f62Olv44VZ(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumxFQvN2P6UDmj430f62Olv44VZ
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

