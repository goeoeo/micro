
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumB0mrxetYaLaA9bxkZv2sH6xDV = []byte("B0mrxetYaLaA9bxkZv2sH6xDV")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesB0mrxetYaLaA9bxkZv2sH6xDV(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumB0mrxetYaLaA9bxkZv2sH6xDV
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

