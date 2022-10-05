package common

import "os"

// 判断路径文件或路径文件夹是否存在
func IsExists(path string) bool {
	_, err := os.Stat(path) // 获取路径信息
	return err == nil
}

// 判断路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path) // 获取路径信息
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断路径是否为文件
func IsFile(path string) bool {
	s, err := os.Stat(path) // 获取路径信息
	if err != nil {
		return false
	}
	return !s.IsDir()
}
