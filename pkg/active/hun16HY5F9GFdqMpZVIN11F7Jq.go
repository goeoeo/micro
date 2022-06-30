
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumhun16HY5F9GFdqMpZVIN11F7Jq = []byte("hun16HY5F9GFdqMpZVIN11F7Jq")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateByteshun16HY5F9GFdqMpZVIN11F7Jq(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumhun16HY5F9GFdqMpZVIN11F7Jq
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

