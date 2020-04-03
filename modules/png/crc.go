package png

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"os"
	"strings"
)

func (p pngFile) CheckAllChunks(fix bool) {
	if !p.chunks[0].CRCIsValid() {
		p.CheckFirstChunk(fix)
	}

	fmt.Printf("\nChecking All chunks...\n")
	for _, chunk := range p.chunks {
		fmt.Println(chunk)
		if !chunk.CRCIsValid() && fix {
			p.file.Seek(chunk.CRCOffset(), io.SeekStart)
			binary.Write(p.file, binary.BigEndian, chunk.CalculateCRC())
		}
	}
	_ = p.file.Sync()
}

func (p pngFile) CheckFirstChunk(fix bool) (valid bool) {
	if len(p.chunks) == 0 {
		return false
	}
	chunk := p.chunks[0]
	fmt.Println("\nChecking first chunk...")
	// ==== ==== ==== ====
	//        w    h
	allbytes := chunk.Bytes()
	fmt.Printf("Calculated CRC: 0x%x\n", crc32.ChecksumIEEE(allbytes))
	fmt.Printf("Recorded CRC: 0x%x\n", chunk.CRC)

	if chunk.CRCIsValid() {
		fmt.Println("True CRC")
		return true
	} else {
		valid = false
		fmt.Println("False CRC, maybe height or weight is false")
		fmt.Println("Now Enum Height...")
		height, err := chunk.EnumHeight(4096)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			if fix {
				name := strings.ReplaceAll(p.file.Name(), ".png", "_fix_height.png")
				p.fixPNG(name, 0, uint32(height))
			}
			return
		}
		fmt.Println("Now Enum Weight...")
		weight, err := chunk.EnumWeight(4096)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			if fix {
				name := strings.ReplaceAll(p.file.Name(), ".png", "_fix_weight.png")
				p.fixPNG(name, uint32(weight), 0)
			}
			return
		}
		height, weight, err = chunk.EnumAll(4096)
		if err != nil {
			fmt.Println("Final not found with max:", 4096)
		} else if fix {
			name := strings.ReplaceAll(p.file.Name(), ".png", "_fix.png")
			p.fixPNG(name, uint32(weight), uint32(height))
		}
		return
	}
}

func (c Chunk) EnumWeight(wmax int) (height int, err error) {
	for weight := 0; weight < wmax; weight++ {
		tmpdata := c.Bytes()
		tmpdata[4] = byte(weight >> 24)
		tmpdata[5] = byte(weight >> 16)
		tmpdata[6] = byte(weight >> 8)
		tmpdata[7] = byte(weight)
		tmpcrc := crc32.ChecksumIEEE(tmpdata)
		if tmpcrc == c.CRC {
			fmt.Printf("Real weight found: 0x%x  =  %d\n", weight, weight)
			return
		}
	}
	return 0, errors.New(fmt.Sprintln("weight not found with max:", wmax))
}

func (c Chunk) EnumHeight(hmax int) (height int, err error) {
	for height = 0; height < hmax; height++ {
		tmpdata := c.Bytes()
		tmpdata[8] = byte(height >> 24)
		tmpdata[9] = byte(height >> 16)
		tmpdata[10] = byte(height >> 8)
		tmpdata[11] = byte(height)
		tmpcrc := crc32.ChecksumIEEE(tmpdata)
		if tmpcrc == c.CRC {
			fmt.Printf("Real height found: 0x%x  =  %d\n", height, height)
			return
		}
	}
	return 0, errors.New(fmt.Sprintln("height not found with max:", hmax))
}

func (c Chunk) EnumAll(max int) (height, weight int, err error) {
	tmpdata := c.Bytes()
	for weight = 0; weight < max; weight++ {
		for height = 0; height < max; height++ {
			tmpdata[4] = byte(weight >> 24)
			tmpdata[5] = byte(weight >> 16)
			tmpdata[6] = byte(weight >> 8)
			tmpdata[7] = byte(weight)
			tmpdata[8] = byte(height >> 24)
			tmpdata[9] = byte(height >> 16)
			tmpdata[10] = byte(height >> 8)
			tmpdata[11] = byte(height)
			tmpcrc := crc32.ChecksumIEEE(tmpdata)
			if tmpcrc == c.CRC {
				fmt.Printf("Real weight found: 0x%x  =  %d\n", weight, weight)
				return
			}
		}
	}
	return 0, 0, errors.New(fmt.Sprintln("weight and height not found with max:", max))
}

func (p pngFile) fixPNG(filePath string, weight, height uint32) {
	fmt.Println("\nTry to fix png file...")
	// 打开原始文件
	originalFile, err := os.Open(p.path)
	if err != nil {
	}
	defer originalFile.Close()
	// 创建新的文件作为目标文件
	newFile, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer newFile.Close()
	// 从源中复制字节到目标文件
	_, err = io.Copy(newFile, originalFile)
	if err != nil {
		return
	}

	// FIX
	if weight > 0 {
		w := make([]byte, 4)
		binary.BigEndian.PutUint32(w, weight)
		_, err = newFile.WriteAt(w, 16)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	if height > 0 {
		h := make([]byte, 4)
		binary.BigEndian.PutUint32(h, height)
		_, err = newFile.WriteAt(h, 20)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	// 将文件内容flush到硬盘中
	err = newFile.Sync()
	if err != nil {
	}
	fmt.Println("fix png file succeed")
}
