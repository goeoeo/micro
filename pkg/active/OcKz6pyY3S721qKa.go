
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumOcKz6pyY3S721qKa = []byte("OcKz6pyY3S721qKa")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesOcKz6pyY3S721qKa(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumOcKz6pyY3S721qKa
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

