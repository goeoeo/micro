
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumEK56Qlwi5MBWMEd2YL6h8sXCzEt9I5p = []byte("EK56Qlwi5MBWMEd2YL6h8sXCzEt9I5p")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesEK56Qlwi5MBWMEd2YL6h8sXCzEt9I5p(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumEK56Qlwi5MBWMEd2YL6h8sXCzEt9I5p
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

