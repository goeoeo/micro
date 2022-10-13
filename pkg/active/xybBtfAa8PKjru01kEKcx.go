
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumxybBtfAa8PKjru01kEKcx = []byte("xybBtfAa8PKjru01kEKcx")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesxybBtfAa8PKjru01kEKcx(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumxybBtfAa8PKjru01kEKcx
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

