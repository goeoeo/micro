
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNum5P2Wo9gAy8IAa0Y = []byte("5P2Wo9gAy8IAa0Y")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytes5P2Wo9gAy8IAa0Y(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNum5P2Wo9gAy8IAa0Y
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

