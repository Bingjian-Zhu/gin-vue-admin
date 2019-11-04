package logger

//ILogger 定义日志输出接口
type ILogger interface {
	////init 初始化日志配置
	Init()
	//Info 打印信息
	Info(args ...interface{})
	//Infof 打印信息，附带template信息
	Infof(template string, args ...interface{})
	//Warn 打印警告
	Warn(args ...interface{})
	//Warnf 打印警告，附带template信息
	Warnf(template string, args ...interface{})
	//Error 打印错误
	Error(args ...interface{})
	//Errorf 打印错误，附带template信息
	Errorf(template string, args ...interface{})
	//Panic 打印Panic信息
	Panic(args ...interface{})
	//Panicf 打印Panic信息，附带template信息
	Panicf(template string, args ...interface{})
	//DPanic 打印DPanic信息，附带template信息
	DPanic(args ...interface{})
	//DPanicf 打印DPanic信息
	DPanicf(template string, args ...interface{})
}
