package util

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

// Returns the base 64 SHA256 of given string.
func Sha256base64(str string) string {
	hash := sha256.New()
	io.WriteString(hash, str)
	return hex.EncodeToString(hash.Sum(nil))
}
