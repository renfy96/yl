package ylString

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBuild(t *testing.T) {
	Convey("字符串拼接", t, func() {
		var s string
		newStr := Build(s, "123")
		So(newStr, ShouldEqual, "123")
	})
}

func TestIsEmpty(t *testing.T) {
	Convey("判断字符串是否为空 - 空", t, func() {
		var s string
		res := IsEmpty(s)
		So(res, ShouldEqual, true)
		Convey("判断字符串是否为空 - 不为空", func() {
			s = "123"
			res = IsEmpty(s)
			So(res, ShouldEqual, false)
		})
	})
}

func TestRemovePrefix(t *testing.T) {
	Convey("获取后缀名 - linux", t, func() {
		path := "/home/go/test.txt"
		res := RemovePrefix(path)
		So(res, ShouldEqual, ".txt")
		Convey("获取后缀名 - windows", func() {
			path := "D:\\golang\\test.txt"
			res := RemoveSuffix(path)
			So(res, ShouldEqual, "test")
		})
	})
}

func TestRemoveSuffix(t *testing.T) {
	Convey("获取文件名 - linux", t, func() {
		path := "/home/go/test.txt"
		res := RemoveSuffix(path)
		So(res, ShouldEqual, "test")
		Convey("获取文件名 - windows", func() {
			path := "D:\\golang\\test.txt"
			res := RemoveSuffix(path)
			So(res, ShouldEqual, "test")
		})
	})
}
