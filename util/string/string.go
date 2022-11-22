package ylString

import (
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// Build 拼接字符串
func Build(someStr ...string) string {
	var b strings.Builder
	for _, v := range someStr {
		b.WriteString(v)
	}
	return b.String()
}

// IsEmpty 是否为空字符串
func IsEmpty(s string) bool {
	return len(s) == 0
}

// RemoveSuffix 去除文件后缀取得文件名
func RemoveSuffix(s string) string {
	var filenameWithSuffix string
	var sysType = runtime.GOOS
	switch sysType {
	case "windows":
		filenameWithSuffix = path.Base(filepath.ToSlash(s))
	case "linux":
		filenameWithSuffix = path.Base(s)
	}
	fileSuffix := path.Ext(filenameWithSuffix)
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	return filenameOnly
}

// RemovePrefix 获取文件拓展名
func RemovePrefix(s string) string {
	var filenameWithSuffix string
	var sysType = runtime.GOOS
	switch sysType {
	case "windows":
		filenameWithSuffix = path.Base(filepath.ToSlash(s))
	case "linux":
		filenameWithSuffix = path.Base(s)
	}
	return path.Ext(filenameWithSuffix)
}
