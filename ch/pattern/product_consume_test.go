package pattern

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRun1(t *testing.T) {
	Convey("方式1：生产者采用协程", t, func() {
		run1()
	})
}

func TestRun2(t *testing.T) {
	Convey("方式2:生产者、消费者都采用协程", t, func() {
		run2()
	})
}

func TestRun3(t *testing.T) {
	Convey("方式3:开关写到消费者内部", t, func() {
		run3()
	})
}
