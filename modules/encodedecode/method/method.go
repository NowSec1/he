package method

var method int

const (
	// base64
	base64Encode = iota
	base64Decode

	// rot13
	rot13Encode
	rot13Decode

	// hex 16进制
	hexEncode
	hexDecode

	// ascii hex
	asciihexEncode
	asciihexDecode

	// octal 8进制
	octalEncode
	octalDecode

	// binary 2进制
	binaryEncode
	binaryDecode

	// url
	urlEncode
	urlDecode

	// utf8
	utf8Encode
	utf8Decode

	// html
	htmlEncode
	htmlDecode
)
