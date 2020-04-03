package png

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func (p pngFile) CheckAllChunks() {
	for _, chunk := range p.chunks {
		fmt.Println(chunk)
		//if !chunk.CRCIsValid() {
		//	p.file.Seek(chunk.CRCOffset(), io.SeekStart)
		//	binary.Write(p.file, binary.BigEndian, chunk.CalculateCRC())
		//}
	}
}

func (p pngFile) CheckFirstChunk() {
	if len(p.chunks) == 0 {
		return
	}
	chunk := p.chunks[0]
	hmax := 4096
	wmax := 4096
	fmt.Println("png info...")
	// ==== ==== ==== ====
	//        w    h
	allbytes := chunk.Bytes()
	fmt.Println(len(allbytes))
	fmt.Println(allbytes)
	fmt.Printf("Now CRC: 0x%x\n", crc32.ChecksumIEEE(allbytes))
	fmt.Printf("Recorded CRC: 0x%x\n", chunk.CRC)

	for h := 0; h < hmax; h++ {
		tmpdata := allbytes
		tmpdata[8] = byte(h >> 24)
		tmpdata[9] = byte(h >> 16)
		tmpdata[10] = byte(h >> 8)
		tmpdata[11] = byte(h)
		tmpcrc := crc32.ChecksumIEEE(tmpdata)
		if tmpcrc == chunk.CRC {
			fmt.Printf("Real height found: 0x%x  =  %d\n", h, h)
			fixHeight("test.png", uint32(h))
			break
		}
	}

	for w := 0; w < wmax; w++ {
		tmpdata := allbytes
		tmpdata[4] = byte(w >> 24)
		tmpdata[5] = byte(w >> 16)
		tmpdata[6] = byte(w >> 8)
		tmpdata[7] = byte(w)
		tmpcrc := crc32.ChecksumIEEE(tmpdata)
		if tmpcrc == chunk.CRC {
			fmt.Printf("Real weight found: 0x%x  =  %d\n", w, w)
			break
		}
	}
}

func fixHeight(filePath string, height uint32) {
	// 打开原始文件
	originalFile, err := os.Open("test.png")
	if err != nil {
	}
	defer originalFile.Close()
	// 创建新的文件作为目标文件
	newFile, err := os.Create("fix.png")
	if err != nil {
	}
	defer newFile.Close()
	// 从源中复制字节到目标文件
	_, err = io.Copy(newFile, originalFile)
	if err != nil {
	}

	// FIX
	h := make([]byte, 4)
	binary.BigEndian.PutUint32(h, height)
	fmt.Println(h)
	_, err = newFile.WriteAt(h, 20)
	if err != nil {
		fmt.Println(err.Error())
	}

	// 将文件内容flush到硬盘中
	err = newFile.Sync()
	if err != nil {
	}
}
