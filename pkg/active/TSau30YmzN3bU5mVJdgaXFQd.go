package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumTSau30YmzN3bU5mVJdgaXFQd = []byte("TSau30YmzN3bU5mVJdgaXFQd")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesTSau30YmzN3bU5mVJdgaXFQd(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumTSau30YmzN3bU5mVJdgaXFQd
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
