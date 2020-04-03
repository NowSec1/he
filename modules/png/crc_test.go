package png

import (
	"testing"
)

func TestExploit(t *testing.T) {
	png, _ := NewPNG("test.png")
	//png.CheckFirstChunk()
	png.CheckAllChunks()
}
