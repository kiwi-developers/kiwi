package utils

import (
	"errors"
	"fmt"
	"os"
)

// search a string key from a slice
// 如果返回值为负数，表明该slice中不存在该该值
// 从两端往slice中间进行搜索，加快搜索速度
func SliceSearchString(s []string, key string) (rs int) {
	var low, high int
	low = 0
	high = len(s) - 1
	for low <= high {
		if s[low] == key {
			return low
		} else {
			low += 1
		}
		if s[high] == key {
			return high
		} else {
			high -= 1
		}
	}
	return -1
}

func DirIsExist(path string) (exist bool, isDir bool) {
	if s, err := os.Stat(path); err == nil {
		if s.IsDir() {
			return true, true
		} else {
			return true, false
		}
	} else {
		return false, false
	}
}

func CheckOrCreateDir(path string) error {
	exist, isDir := DirIsExist(path)
	if exist {
		if isDir {
			return nil
		} else {
			return errors.New(fmt.Sprintf("path %s is existed, but it's not a directory", path))
		}
	} else {
		return os.Mkdir(path, 0644)
	}
}
