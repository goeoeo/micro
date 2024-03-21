
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumVWtrI2Y9Ry5zY4KZv1fo5ZGkmBp1qc = []byte("VWtrI2Y9Ry5zY4KZv1fo5ZGkmBp1qc")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesVWtrI2Y9Ry5zY4KZv1fo5ZGkmBp1qc(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumVWtrI2Y9Ry5zY4KZv1fo5ZGkmBp1qc
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

