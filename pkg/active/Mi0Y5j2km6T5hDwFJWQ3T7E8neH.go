
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumMi0Y5j2km6T5hDwFJWQ3T7E8neH = []byte("Mi0Y5j2km6T5hDwFJWQ3T7E8neH")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesMi0Y5j2km6T5hDwFJWQ3T7E8neH(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumMi0Y5j2km6T5hDwFJWQ3T7E8neH
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

