package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumLEpP5eVlGW = []byte("LEpP5eVlGW")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesLEpP5eVlGW(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumLEpP5eVlGW
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
