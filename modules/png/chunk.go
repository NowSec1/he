package png

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

const chunkStartOffset = 8
const endChunk = "IEND"

type Chunk struct {
	Offset        int64
	Length        uint32
	Type          [4]byte
	Data          []byte
	CRC           uint32
	CalculatedCRC uint32
}

func readChunk(reader *os.File) (chunk Chunk, err error) {
	chunk.Offset, err = reader.Seek(0, io.SeekCurrent)
	if err != nil {
		return
	}
	// chunk Length 4Bytes
	err = binary.Read(reader, binary.BigEndian, &chunk.Length)
	if err != nil {
		return
	}

	// chunk Type 4Bytes
	err = binary.Read(reader, binary.BigEndian, &chunk.Type)
	if err != nil {
		return
	}

	// chunk Data var
	chunk.Data = make([]byte, chunk.Length)
	read, err := reader.Read(chunk.Data)
	if read == 0 || err != nil {
		return
	}

	// chunk CRC 4Bytes
	err = binary.Read(reader, binary.BigEndian, &chunk.CRC)
	if err != nil {
		return
	}

	// calculate real CRC (maybe NOT the correct one)
	chunk.CalculatedCRC = chunk.CalculateCRC()
	return chunk, nil
}

func (c Chunk) String() string {
	return fmt.Sprintf("%s@%x - %X - Valid CRC? %v", c.Type, c.Offset, c.CRC, c.CRCIsValid())
}

func (c Chunk) Bytes() []byte {
	var buffer bytes.Buffer

	binary.Write(&buffer, binary.BigEndian, c.Type)
	buffer.Write(c.Data)

	return buffer.Bytes()
}

func (c Chunk) CRCIsValid() bool {
	if string(c.Type[:]) == endChunk {
		return true
	}
	return c.CRC == c.CalculateCRC()
}

func (c Chunk) CalculateCRC() uint32 {
	if c.CalculatedCRC != 0 {
		return c.CalculatedCRC
	}
	return crc32.ChecksumIEEE(c.Bytes())
}

func (c Chunk) CRCOffset() int64 {
	return c.Offset + int64(8+c.Length)
}
