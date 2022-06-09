package crypto

import (
	"crypto/aes"
	"crypto/cipher"
)

func blockUnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:length-unpadding]
}

func Decrypt(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	decryptData := make([]byte, len(data))
	mode.CryptBlocks(decryptData, data)

	return blockUnPadding(decryptData), nil
}
