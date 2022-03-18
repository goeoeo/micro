package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumWWeaalGIN7jR3tmWvd = []byte("WWeaalGIN7jR3tmWvd")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesWWeaalGIN7jR3tmWvd(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumWWeaalGIN7jR3tmWvd
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
