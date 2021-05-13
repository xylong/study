package config

// ErrorChan 错误消息
// 存储启动时候产生的错误
var ErrorChan chan error

func init() {
	ErrorChan = make(chan error)
}
