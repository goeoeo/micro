package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumT5Db2dCCI0txNPLdYPY8Y3gUFo = []byte("T5Db2dCCI0txNPLdYPY8Y3gUFo")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesT5Db2dCCI0txNPLdYPY8Y3gUFo(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumT5Db2dCCI0txNPLdYPY8Y3gUFo
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
