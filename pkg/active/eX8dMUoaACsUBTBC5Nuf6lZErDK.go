
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumeX8dMUoaACsUBTBC5Nuf6lZErDK = []byte("eX8dMUoaACsUBTBC5Nuf6lZErDK")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateByteseX8dMUoaACsUBTBC5Nuf6lZErDK(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumeX8dMUoaACsUBTBC5Nuf6lZErDK
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

