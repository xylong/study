package pack

import (
	"io"
	"log"
	"os"
)

var (
	Trace   *log.Logger // 记录一切
	Info    *log.Logger // 记录重要信息
	Warning *log.Logger // 记录警告
	Error   *log.Logger // 记录错误
)

func TraceLog() {
	// 丢弃日志
	Trace = log.New(io.Discard, "Trace: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func InfoLog() {
	Info = log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func WarningLog() {
	Warning = log.New(os.Stdout, "Warning: ", log.Ldate|log.Ltime|log.Llongfile)
}

func ErrorLog() {
	file, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Panic("日志文件打开失败")
	}
	Error = log.New(file, "Error: ", log.Ldate|log.Lmicroseconds|log.Llongfile)
}
