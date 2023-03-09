
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNum16Q1XBBLEroh1DTK2q = []byte("16Q1XBBLEroh1DTK2q")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytes16Q1XBBLEroh1DTK2q(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNum16Q1XBBLEroh1DTK2q
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

