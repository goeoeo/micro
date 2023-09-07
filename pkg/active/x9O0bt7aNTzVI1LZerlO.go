
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumx9O0bt7aNTzVI1LZerlO = []byte("x9O0bt7aNTzVI1LZerlO")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesx9O0bt7aNTzVI1LZerlO(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumx9O0bt7aNTzVI1LZerlO
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

