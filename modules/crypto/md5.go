package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Hash(data string) (result string, err error) {
	has := md5.Sum([]byte(data))
	result = hex.EncodeToString(has[:])
	return
}
