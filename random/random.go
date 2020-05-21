package random

import (
	"crypto/rand"
	"io"
)

// NewRandom returns length random bytes
// length is the length of the byte array which will be returned
// Return length random bytes and an error if retrieving the random bytes returns an error.
func NewRandom(length int) (random []byte, err error) {
	random = make([]byte, length)
	_, err = io.ReadFull(rand.Reader, random[:])
	if err != nil {
		return nil, err
	}
	return
}