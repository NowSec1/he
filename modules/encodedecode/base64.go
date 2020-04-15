package encodedecode

import (
	b64 "encoding/base64"
)

func Base64StdEncode(data string) (result string, err error) {
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	return sEnc, nil
}

func Base64StdDecodeString(data string) (result string, err error) {
	sDec, err := b64.StdEncoding.DecodeString(data)
	if err != nil {
		return
	}
	return string(sDec), nil
}

func Base64URLEncode(data string) (result string, err error) {
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	return uEnc, nil
}

func Base64URLDecode(data string) (result string, err error) {
	uDec, err := b64.URLEncoding.DecodeString(data)
	if err != nil {
		return
	}
	return string(uDec), nil
}
