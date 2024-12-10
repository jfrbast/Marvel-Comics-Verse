package requests

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Hash(ts int, publicKey []string, privateKey []string) string {
	token := fmt.Sprintf("%d%s%s", ts, privateKey, publicKey)
	h := md5.New()
	h.Write([]byte(token))
	return hex.EncodeToString(h.Sum(nil))
}
