
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNum2mJTnWxn0aT5IEithz = []byte("2mJTnWxn0aT5IEithz")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytes2mJTnWxn0aT5IEithz(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNum2mJTnWxn0aT5IEithz
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

