package examples

import "go.uber.org/zap/zapcore"

type Other struct {
	Val int
}

func (l *Other) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt("val", l.Val)
	return nil
}
