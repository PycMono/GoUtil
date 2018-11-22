package fileUtil

import (
	"archive/zip"
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
)

// zip压缩文件（不支持目录压缩）
// 参数：
// sourceList：需要压缩的文件路径列表
// targetPath：压缩到目标路径
// 返回值：
// 1.错误对象
func Zip(sourceList []string, targetPath string) error {
	targetFile, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	newZip := zip.NewWriter(targetFile)
	defer newZip.Close()

	// 对源文件目录遍历处理，获取文件信息
	for _, item := range sourceList {
		// 获取文件信息
		info, err := os.Stat(item)
		if err != nil {
			return err
		}
		if info.IsDir() {
			continue
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = filepath.Base(item)
		header.Method = zip.Deflate // 默认是Store，只归档文件不压缩，Deflate：压缩保存

		var writer io.Writer
		if writer, err = newZip.CreateHeader(header); err != nil {
			return err
		}

		file, err := os.Open(item)
		if err != nil {
			return err
		}
		defer file.Close()

		if _, err = io.Copy(writer, file); err != nil {
			return err
		}
	}

	return nil
}

// 解压文件()
// 参数：
// sourceFilePath：原文件夹路径
// tarFilePath：解压到目标文件夹路径
// 返回值：
// 1.错误对象
func UnZip(sourceFilePath, tarFilePath string) error {
	// 判断目标文件夹是否存在，如果不存在就创建
	if !DirExists(tarFilePath) {
		err := os.MkdirAll(tarFilePath, os.ModePerm|os.ModeTemporary)
		if err != nil {
			return err
		}
	}

	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	zipFile, err := zip.OpenReader(sourceFile.Name())
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// 遍历文件处理数据
	for _, file := range zipFile.File {
		tempTarFilePath := path.Join(tarFilePath, file.Name)
		fmt.Println(tempTarFilePath)
		fileInfo := file.FileInfo()
		if fileInfo.IsDir() {
			err = os.MkdirAll(tempTarFilePath, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}

		srcFile, err := file.Open()
		if err != nil {
			return err
		}
		defer srcFile.Close()

		tarFile, err := os.Create(tempTarFilePath)
		if err != nil {
			return err
		}
		defer tarFile.Close()

		io.Copy(tarFile, srcFile)
	}

	return nil
}

func Zip2(data []byte, level int) ([]byte, error) {
	var buffer bytes.Buffer
	compressor, err := zlib.NewWriterLevelDict(&buffer, level, nil)
	if err != nil {
		return nil, err
	}

	compressor.Write(data)
	compressor.Close()

	return buffer.Bytes(), nil
}
