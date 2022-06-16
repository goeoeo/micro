
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumbsQalhkPnocAV0wLVvD5NYsKKS = []byte("bsQalhkPnocAV0wLVvD5NYsKKS")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesbsQalhkPnocAV0wLVvD5NYsKKS(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumbsQalhkPnocAV0wLVvD5NYsKKS
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

