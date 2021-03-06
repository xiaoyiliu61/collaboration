package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

func MD5HashString(data string) string {
	md5Hash:=md5.New()
	md5Hash.Write([]byte(data))
	bytes:=md5Hash.Sum(nil)
	return hex.EncodeToString(bytes)
}
func Base64Str(msg string) string {
	return base64.StdEncoding.EncodeToString([]byte(msg))
}
