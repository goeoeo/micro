
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumM3DuvCR9iGns = []byte("M3DuvCR9iGns")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesM3DuvCR9iGns(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumM3DuvCR9iGns
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

