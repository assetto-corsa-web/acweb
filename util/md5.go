package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

// Returns the base 64 MD5 of given string.
func Md5base64(str string) string {
	hash := md5.New()
	io.WriteString(hash, str)
	return hex.EncodeToString(hash.Sum(nil))
}
