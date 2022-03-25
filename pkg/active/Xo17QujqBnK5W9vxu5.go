package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumXo17QujqBnK5W9vxu5 = []byte("Xo17QujqBnK5W9vxu5")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesXo17QujqBnK5W9vxu5(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumXo17QujqBnK5W9vxu5
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
