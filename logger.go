package mongo

type Logger interface {
	Log(args ...interface{})
	Logf(pattern string, args ...interface{})
	Error(args ...interface{})
	Errorf(patter string, args ...interface{})
}
