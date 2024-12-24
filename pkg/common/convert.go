package common

import (
	"crypto/sha1"
	"encoding/hex"
)

func ConvertPassword(password string) string {
	h := sha1.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}
