package log

type GRPCLogger struct {
	*RollFileLogger
}

func NewGRPCLogger() *GRPCLogger {
	return &GRPCLogger{RollFileLogger: lg}
}

func (*GRPCLogger) Info(args ...interface{})                    {}
func (*GRPCLogger) Infoln(args ...interface{})                  {}
func (*GRPCLogger) Infof(format string, args ...interface{})    {}
func (*GRPCLogger) Warning(args ...interface{})                 {}
func (*GRPCLogger) Warningln(args ...interface{})               {}
func (*GRPCLogger) Warningf(format string, args ...interface{}) {}
func (*GRPCLogger) V(l int) bool {
	return false
}
