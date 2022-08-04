
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumJ1UpuTka0ET01Pj = []byte("J1UpuTka0ET01Pj")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesJ1UpuTka0ET01Pj(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumJ1UpuTka0ET01Pj
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

