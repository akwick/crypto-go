// The code example for encryption is based upon the encryption
// presented by George Tankersley <george.tankersley@gmail.com>
// <https://github.com/gtank/cryptopasta/blob/master/encrypt.go>
// under the CC0 Public Domain Dedication license
// <http://creativecommons.org/publicdomain/zero/1.0/>.
// The code examples are slightly adapted by Anna-Katharina Wickert
// under the same license.
package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

var key [32] byte
var keyset bool

// GetKey returns a random key for AES-256.
// panics if the key couldn't be set.
func GetKey() [32]byte {
	if !keyset {
		err := setKey()
		if err != nil {
			panic(err)
		}
	}
	return key
}

// Encrypt encrypts a plaintext using AES-256-GCM.
// GCM in combination with AES not only hides the plaintext
// but also check if the message was changed.
// Returns the output in the form nonce|ciphertext|tag with '|'
// presenting a concatenation.
func Encrypt(plaintext []byte, key *[32]byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func Decrypt(ciphertext []byte, key *[32]byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("malformed ciphertext")
	}

	return gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil)
}


// setKey generates a random key for AES-256.
// returns an error if a random number couldn't created.
func setKey() error{
	_, err := io.ReadFull(rand.Reader, key[:])
	if err != nil {
		return err
	}
	keyset = true
	return nil
}