
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumB2H8mbtybY3XNo = []byte("B2H8mbtybY3XNo")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesB2H8mbtybY3XNo(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumB2H8mbtybY3XNo
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

