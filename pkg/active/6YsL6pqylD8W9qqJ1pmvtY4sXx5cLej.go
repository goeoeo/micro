
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNum6YsL6pqylD8W9qqJ1pmvtY4sXx5cLej = []byte("6YsL6pqylD8W9qqJ1pmvtY4sXx5cLej")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytes6YsL6pqylD8W9qqJ1pmvtY4sXx5cLej(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNum6YsL6pqylD8W9qqJ1pmvtY4sXx5cLej
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

