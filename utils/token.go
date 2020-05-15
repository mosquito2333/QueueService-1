package utils

import (
	"crypto/md5"
	cryp "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Token字串
func GetToken() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(cryp.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}
