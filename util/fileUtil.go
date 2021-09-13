package util

import (
	"runtime"
	"strings"
)

// 获取项目运行路径，即 /main.go 的位置
func GetRunPath() (filePathStr string) {
	// 第一次用 go 正式写代码，所以可能走了弯路
	_, fileStr, _, _ := runtime.Caller(1)
	// fileStr : ....../parsePatentWord/read/readTry.go
	// 所以要获得根目录路径，就把 /read/readTry.go 忽略
	filenameSlice := strings.Split(fileStr, "/")
	filePathSlice := filenameSlice[:len(filenameSlice)-2]
	filePathStr = strings.Join(filePathSlice, "\\")
	return
}
