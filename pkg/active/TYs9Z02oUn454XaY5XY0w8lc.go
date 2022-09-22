
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumTYs9Z02oUn454XaY5XY0w8lc = []byte("TYs9Z02oUn454XaY5XY0w8lc")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesTYs9Z02oUn454XaY5XY0w8lc(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumTYs9Z02oUn454XaY5XY0w8lc
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

