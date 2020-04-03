package png

import (
	"fmt"
	"io"
	"os"
)

type pngFile struct {
	file   *os.File
	path   string
	chunks []Chunk
}

func NewPNG(filePath string) (png pngFile, err error) {
	fmt.Println("Path:", filePath)
	png.path = filePath
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	png.file = file
	defer file.Close()

	fileInfo, err := file.Stat()
	fmt.Printf("Name: %s\nSize: %d B\nModTime: %s\n", fileInfo.Name(), fileInfo.Size(), fileInfo.ModTime())

	if !isPng(file) {
		fmt.Fprintln(os.Stderr, "Not a PNG")
		os.Exit(-1)
	}

	// Read all the chunks. They start with IHDR at offset 8
	png.chunks, err = png.ReadAllChunks()
	if err != nil {
		return
	}
	return
}

func (p pngFile) ReadAllChunks() (chunks []Chunk, err error) {
	reader := p.file
	_, err = reader.Seek(chunkStartOffset, io.SeekStart)
	if err != nil {
		return
	}

	for {
		var chunk Chunk
		chunk, err = readChunk(reader)
		if err != nil {
			break
		}
		chunks = append(chunks, chunk)
		if string(chunk.Type[:]) == endChunk {
			return
		}
	}
	return
}

func isPng(f *os.File) bool {
	_, err := f.Seek(1, io.SeekStart)
	if err != nil {
		return false
	}

	magic := make([]byte, 3)
	_, err = f.Read(magic)
	if err != nil {
		return false
	}

	return string(magic) == "PNG"
}
