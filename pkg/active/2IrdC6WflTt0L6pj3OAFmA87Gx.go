
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNum2IrdC6WflTt0L6pj3OAFmA87Gx = []byte("2IrdC6WflTt0L6pj3OAFmA87Gx")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytes2IrdC6WflTt0L6pj3OAFmA87Gx(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNum2IrdC6WflTt0L6pj3OAFmA87Gx
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

