package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha512"

	"golang.org/x/crypto/pbkdf2"
)

func GenerateKeyandIv(initKey []byte) ([]byte, []byte) {
	nKey := 32
	nIv := 16
	salt := []byte("_@<$>^?*")

	key := pbkdf2.Key(initKey, salt, 4096, nKey+nIv, sha512.New)

	return key[:nKey], key[nKey:]
}

func BlockPadding(text []byte, blockSize int) []byte {
	if len(text)%blockSize != 0 {
		paddingSize := blockSize - len(text)%blockSize
		paddingBytes := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)

		text = append(text, paddingBytes...)
	}

	return text
}

func Encrypt(data []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)

	cipherText := make([]byte, len(data))
	mode.CryptBlocks(cipherText, data)

	return cipherText, nil
}
