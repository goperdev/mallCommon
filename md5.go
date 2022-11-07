package mallCommon

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5Encoding md5编码
func Md5Encoding(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
