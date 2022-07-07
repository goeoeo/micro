
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumT3r5ENU90pDNofcFtt741XlP7ExQ = []byte("T3r5ENU90pDNofcFtt741XlP7ExQ")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesT3r5ENU90pDNofcFtt741XlP7ExQ(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumT3r5ENU90pDNofcFtt741XlP7ExQ
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

