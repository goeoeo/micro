
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNum82vNsbZ92QARUvz4pgFu0DRrwdN = []byte("82vNsbZ92QARUvz4pgFu0DRrwdN")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytes82vNsbZ92QARUvz4pgFu0DRrwdN(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNum82vNsbZ92QARUvz4pgFu0DRrwdN
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

