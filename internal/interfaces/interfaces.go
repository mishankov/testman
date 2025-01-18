package interfaces

type TB interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Helper()
}
