package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

//FileExist 判断文件是否存在
func FileExist(filePath string) (bool, error) {
	var file *os.File
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()
	return true, nil
}

//DirExist 判断文件夹是否存在
func DirExist(dir string) bool {
	_, err := os.Stat(dir)
	return !os.IsNotExist(err)
}

//CreateDir 创建文件夹
func CreateDir(dir string) error {
	err := os.Mkdir(dir, os.ModePerm)
	return err
}

//DeleteFile 删除文件
func DeleteFile(filePath string) error {
	exist, _ := FileExist(filePath)
	if exist {
		return os.Remove(filePath)
	}
	return nil
}

//GetFileSize 获取文件大小
func GetFileSize(destPath string) (int64, error) {
	file, err := os.Open(destPath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return 0, err
	}
	size := fi.Size()
	return size, nil
}


//GetFileMD5String 获取文件md5值
func GetFileMD5String(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return ""
	}

	md5 := md5.New()
	io.Copy(md5, file)
	lastString := hex.EncodeToString(md5.Sum(nil))
	return lastString
}
