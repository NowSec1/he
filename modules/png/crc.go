package png

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
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

func (c Chunk) EnumWeight(wmax int) (weight int, err error) {
	tmpdata := c.Bytes()
	for weight = 0; weight < wmax; weight++ {
		binary.BigEndian.PutUint32(tmpdata[4:8], uint32(weight))
		tmpcrc := crc32.ChecksumIEEE(tmpdata)
		if tmpcrc == c.CRC {
			fmt.Printf("Real weight found: 0x%x  =  %d\n", weight, weight)
			return
		}
	}
	return 0, errors.New(fmt.Sprintln("weight not found with max:", wmax))
}

func (c Chunk) EnumHeight(hmax int) (height int, err error) {
	tmpdata := c.Bytes()
	for height = 0; height < hmax; height++ {
		binary.BigEndian.PutUint32(tmpdata[8:12], uint32(height))
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
			binary.BigEndian.PutUint32(tmpdata[4:8], uint32(weight))
			binary.BigEndian.PutUint32(tmpdata[8:12], uint32(height))
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
	originData, err := ioutil.ReadFile(p.path)
	if err != nil {
		return
	}

	// FIX
	if weight > 0 {
		binary.BigEndian.PutUint32(originData[16:20], weight)
	}

	if height > 0 {
		binary.BigEndian.PutUint32(originData[20:24], height)
	}

	err = ioutil.WriteFile(filePath, originData, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("fix png file succeed")
}
