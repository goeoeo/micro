
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumeAGo9vWUUGD6Q = []byte("eAGo9vWUUGD6Q")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateByteseAGo9vWUUGD6Q(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumeAGo9vWUUGD6Q
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

