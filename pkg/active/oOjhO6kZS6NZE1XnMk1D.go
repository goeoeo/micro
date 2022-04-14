package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumoOjhO6kZS6NZE1XnMk1D = []byte("oOjhO6kZS6NZE1XnMk1D")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesoOjhO6kZS6NZE1XnMk1D(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumoOjhO6kZS6NZE1XnMk1D
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
