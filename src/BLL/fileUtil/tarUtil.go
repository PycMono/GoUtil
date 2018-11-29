package fileUtil

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
)

// Tar压缩(不支持压缩目录哦)
// 参数：
// sourceList：需要压缩的文件路径列表
//targetPath：压缩到目标路径
// 返回值：
// 1.错误对象
func Tar(sourceList []string, targetPath string) error {
	targetFile, err := os.Create(targetPath)
	if err != nil {
		return err
	}

	gzipWriter := gzip.NewWriter(targetFile)
	newWriter := tar.NewWriter(gzipWriter)

	defer func() {
		targetFile.Close()
		gzipWriter.Close()
		newWriter.Close()
	}()

	// 遍历源文件处理
	for _, item := range sourceList {
		// 获取文件信息
		fileInfo, err := os.Stat(item)
		if err != nil {
			// 这里最好记录当前未找到的路径，不要直接返回，把后面想压缩的文件都给抛弃了
			return err
		}

		if fileInfo.IsDir() {
			continue
		}

		// 判断文件是否是标准文件
		if !fileInfo.Mode().IsRegular() {
			return nil
		}

		header, err := tar.FileInfoHeader(fileInfo, fileInfo.Name())
		if err != nil {
			return err
		}

		header.Name = filepath.Base(item)

		if err := newWriter.WriteHeader(header); err != nil {
			return err
		}

		file, err := os.Open(item)
		if err != nil {
			return err
		}
		defer file.Close()

		if _, err = io.Copy(newWriter, file); err != nil {
			return err
		}
	}

	return nil
}

// 解压文件
// 参数：
// sourceFilePath：原文件夹路径
// tarFilePath：解压到目标文件夹路径
// 返回值：
// 1.错误对象
func UnTar(sourceFilePath, tarFilePath string) error {
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

	newGzip, err := gzip.NewReader(sourceFile)
	if err != nil {
		return err
	}

	defer func() {
		sourceFile.Close()
		newGzip.Close()
	}()

	newTar := tar.NewReader(newGzip)
	for {
		header, err := newTar.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}

		targetFilePath := filepath.Join(tarFilePath, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(targetFilePath, info.Mode()); err != nil {
				return err
			}
			continue
		}

		tempTarFilePath := path.Join(tarFilePath, header.Name)
		file, err := os.OpenFile(tempTarFilePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if err != nil {
			fmt.Println(1)
			return err
		}
		defer file.Close()

		_, err = io.Copy(file, newTar)
		if err != nil {
			fmt.Println(2)
			return err
		}
	}

	return nil
}
