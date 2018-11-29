package fileUtil

import (
	"compress/zlib"
	"fmt"
	"testing"
)

func TestZip(t *testing.T) {
	data := ([]byte)(InitString)
	result, _ := Zlib(data, zlib.DefaultCompression)
	ZlibBytes = result

	data, _ = UnZlib(ZlibBytes)
	result2 := string(data)
	if result2 != InitString {
		fmt.Println(fmt.Printf("解压缩失败，源数据%s，实际%s", InitString, result2))
	}

	fmt.Print(result2)
}
