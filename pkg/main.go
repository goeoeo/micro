package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	r "math/rand"
	"os"
	"time"
)

var funBody = `
package active

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNum%s = []byte("%s")

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytes%s(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNum%s
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
			bytes[i] = alphabets[b%%byte(len(alphabets))]
		}
	}
	return bytes
}

`

func main() {
	r.Seed(time.Now().UnixNano())

	name := string(RandomCreateBytes(r.Intn(32)))
	funBody = fmt.Sprintf(funBody, name, name, name, name)

	fileName := fmt.Sprintf("./active/%s.go", name)
	err := ioutil.WriteFile(fileName, []byte(funBody), os.ModePerm)
	if err != nil {
		log.Fatalf("ioutil.WriteFile: err = %v", err)
	}
}

var alphaNum = []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytes(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNum
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
