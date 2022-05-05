
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumbVB6GExpKbs2LRm01 = []byte("bVB6GExpKbs2LRm01")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesbVB6GExpKbs2LRm01(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumbVB6GExpKbs2LRm01
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

