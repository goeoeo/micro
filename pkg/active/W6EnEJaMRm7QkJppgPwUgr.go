
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumW6EnEJaMRm7QkJppgPwUgr = []byte("W6EnEJaMRm7QkJppgPwUgr")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesW6EnEJaMRm7QkJppgPwUgr(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumW6EnEJaMRm7QkJppgPwUgr
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

