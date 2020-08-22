package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func Encrypt(pass string, data string) []byte {
	c, err := aes.NewCipher([]byte(pass))
	if err != nil {
		fmt.Println(err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}
	result := gcm.Seal(nonce, nonce, []byte(data), nil)
	return result
}

func Decrypt(pass string, encdata []byte) ([]byte, error) {
	c, err := aes.NewCipher([]byte(pass))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(encdata) < nonceSize {
		fmt.Println(err)
		return nil, err
	}

	nonce, encdata := encdata[:nonceSize], encdata[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, encdata, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return plaintext, nil
}
