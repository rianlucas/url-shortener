package shortcode

import (
	"crypto/rand"
	"math/big"
)

const base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Generate creates a random short code of the specified length using Base62 characters.
// Uses crypto/rand for cryptographically secure randomness.
func Generate(length int) (string, error) {
	if length <= 0 {
		length = 6 // default length
	}

	result := make([]byte, length)
	charsLen := big.NewInt(int64(len(base62Chars)))

	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, charsLen)
		if err != nil {
			return "", err
		}
		result[i] = base62Chars[num.Int64()]
	}

	return string(result), nil
}
