package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
)

// aes encryot with PKCS5Padding and return base64 encoded string
func AESBase64Encrypt(clear, key, iv string) (encryoted string, err error) {
	var block cipher.Block
	if block, err = aes.NewCipher([]byte(key)); err != nil {
		log.Println(err)
		return
	}
	encrypt := cipher.NewCBCEncrypter(block, []byte(iv))
	var source []byte = PKCS5Padding([]byte(clear), block.BlockSize())
	var dst []byte = make([]byte, len(source))
	encrypt.CryptBlocks(dst, source)
	encryoted = base64.StdEncoding.EncodeToString(dst)
	return
}

// aes decryot decode base64 encoded, encrypted data
func AESBase64Decrypt(encrypted, key, iv string) (clear string, err error) {
	var block cipher.Block
	if block, err = aes.NewCipher([]byte(key)); err != nil {
		log.Println(err)
		return
	}
	encrypter := cipher.NewCBCDecrypter(block, []byte(iv))

	var source []byte
	if source, err = base64.StdEncoding.DecodeString(encrypted); err != nil {
		log.Println(err)
		return
	}
	var dst []byte = make([]byte, len(source))
	encrypter.CryptBlocks(dst, source)
	clear = string(PKCS5Unpadding(dst))
	return
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
