package fileUtil

import (
	"bytes"
	"compress/zlib"
	"io"
)

// 字节压缩（zlib方式压缩）
// 参数：
// data：待压缩的数组
// level：等级
// 返回值：
// 1.压缩后的数据
// 2.错误对象
func Zlib(data []byte, level int) ([]byte, error) {
	var buffer bytes.Buffer
	zlibWriter, err := zlib.NewWriterLevelDict(&buffer, level, nil)
	if err != nil {
		return nil, err
	}

	zlibWriter.Write(data)
	zlibWriter.Close()

	return buffer.Bytes(), nil
}

// 字节解压（zlib方式压缩）
// data：待解压的数据
// 返回值:
// 1.解压后的数据
// 2.错误对象
func UnZlib(data []byte) ([]byte, error) {
	dataReader := bytes.NewReader(data)
	zlibReader, err := zlib.NewReader(dataReader)
	if err != nil {
		return nil, err
	}
	defer zlibReader.Close()

	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, zlibReader)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
