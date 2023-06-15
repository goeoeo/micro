
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumem3Z9SPpL8DlLPof = []byte("em3Z9SPpL8DlLPof")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesem3Z9SPpL8DlLPof(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumem3Z9SPpL8DlLPof
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

