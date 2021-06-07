package pack

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTraceLog(t *testing.T) {
	Convey("trace", t, func() {
		TraceLog()
		Trace.Println("foo")
	})
}

func TestInfoLog(t *testing.T) {
	Convey("info", t, func() {
		InfoLog()
		Info.Println("记录重要信息")
	})
}

func TestWarningLog(t *testing.T) {
	Convey("warning", t, func() {
		WarningLog()
		Warning.Println("记录警告信息")
	})
}

func TestErrorLog(t *testing.T) {
	Convey("error", t, func() {
		ErrorLog()
		Error.Println("记录错误")
		s := []struct {
			name string
			age  int
		}{
			{"静静", 20},
			{"琳琳", 18},
		}
		Error.Println(s)
		m := map[string][]struct {
			name string
			age  uint
		}{
			"a": []struct {
				name string
				age  uint
			}{
				{"张三", 13},
				{"李四", 14},
			},
			"b": []struct {
				name string
				age  uint
			}{
				{"静静", 20},
				{"琳琳", 18},
			},
		}
		Error.Println(m)
	})
}
