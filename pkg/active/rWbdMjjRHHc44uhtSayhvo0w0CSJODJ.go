
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNumrWbdMjjRHHc44uhtSayhvo0w0CSJODJ = []byte("rWbdMjjRHHc44uhtSayhvo0w0CSJODJ")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytesrWbdMjjRHHc44uhtSayhvo0w0CSJODJ(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNumrWbdMjjRHHc44uhtSayhvo0w0CSJODJ
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

