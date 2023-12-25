package helpers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"math/big"
	"math/rand"
	"os"
)

func GeneratePhoneNumber() string {
	prefix := "08"

	var randomDigits string
	for i := 0; i < 23; i++ {
		randomDigits += fmt.Sprintf("%d", rand.Intn(10))
	}

	phoneNumber := fmt.Sprintf("%s%s", prefix, randomDigits)
	return phoneNumber
}

func GetGanjilGenapNumber(data []string) (map[string][]string, error) {
	result := make(map[string][]string)

	for _, strNum := range data {
		num, success := new(big.Int).SetString(strNum, 10)
		if !success {
			return nil, fmt.Errorf("failed to parse the number: %s", strNum)
		}

		category := "ganjil"
		if num.Mod(num, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
			category = "genap"
		}

		result[category] = append(result[category], strNum)
	}

	return result, nil
}

func EncryptAESCBC(data string) string {
	key := []byte(os.Getenv("KEY"))
	iv := []byte(os.Getenv("IV"))

	ciphertext := encrypt([]byte(data), key, iv)
	encryptedData := append(iv, ciphertext...)

	return base64.StdEncoding.EncodeToString(encryptedData)
}

func encrypt(plaintext, key, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	plaintext = padPKCS7(plaintext, aes.BlockSize)

	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext, plaintext)

	return ciphertext
}

func padPKCS7(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func DecryptAESCBC(encryptedData string) string {
	key := []byte(os.Getenv("KEY"))
	iv := []byte(os.Getenv("IV"))

	decodedData, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		panic(err)
	}

	decryptedText := decrypt(decodedData[aes.BlockSize:], key, iv)
	if err != nil {
		panic(err)
	}

	return decryptedText
}

func decrypt(ciphertext, key, iv []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	plaintext = unpadPKCS7(plaintext)

	return string(plaintext)
}

func unpadPKCS7(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
