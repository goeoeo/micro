
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumGF5lfzR3OaEzZlykwqvKvmjNSnMaP = []byte("GF5lfzR3OaEzZlykwqvKvmjNSnMaP")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesGF5lfzR3OaEzZlykwqvKvmjNSnMaP(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumGF5lfzR3OaEzZlykwqvKvmjNSnMaP
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

