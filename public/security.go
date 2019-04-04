package public

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}
