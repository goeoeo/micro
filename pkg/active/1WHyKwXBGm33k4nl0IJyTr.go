
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNum1WHyKwXBGm33k4nl0IJyTr = []byte("1WHyKwXBGm33k4nl0IJyTr")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytes1WHyKwXBGm33k4nl0IJyTr(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNum1WHyKwXBGm33k4nl0IJyTr
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

