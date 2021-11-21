package random

import (
	"crypto/rand"
	"math/big"
)

const (
	alphabet       = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	alphabetLength = int64(len(alphabet))
)

func Int(min, max int64) (int64, error) {
	number, err := rand.Int(rand.Reader, big.NewInt(max-min+1))
	if err != nil {
		return 0, err
	}
	return min + number.Int64(), nil
}

func Bytes(length int) ([]byte, error) {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		number, err := rand.Int(rand.Reader, big.NewInt(alphabetLength))
		if err != nil {
			return nil, err
		}
		bytes[i] = alphabet[number.Int64()]
	}
	return bytes, nil
}

func String(length int) (string, error) {
	bytes, err := Bytes(length)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}
