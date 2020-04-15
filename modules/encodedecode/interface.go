package encodedecode

type encodedecode interface {
	encode(data []byte) (result []byte, err error)
	dncode(data []byte) (result []byte, err error)
	typeName() (typestr string, methodIdx method)
}
