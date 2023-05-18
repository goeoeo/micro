
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNum9QCsSyO77fsEsE0oMp = []byte("9QCsSyO77fsEsE0oMp")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytes9QCsSyO77fsEsE0oMp(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNum9QCsSyO77fsEsE0oMp
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

