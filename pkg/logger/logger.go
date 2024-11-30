package logger

type Logger interface {
	InitLogger()
	Debug(...interface{})
}
