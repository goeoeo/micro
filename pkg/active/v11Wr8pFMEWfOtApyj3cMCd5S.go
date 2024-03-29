
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumv11Wr8pFMEWfOtApyj3cMCd5S = []byte("v11Wr8pFMEWfOtApyj3cMCd5S")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesv11Wr8pFMEWfOtApyj3cMCd5S(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumv11Wr8pFMEWfOtApyj3cMCd5S
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

