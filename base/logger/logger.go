package logger

import (
	"gonet/base"
)

func Debug(v ...interface{}) {
	base.GLOG.Debug(v...)
}
func Debugf(f string, v ...interface{}) {
	base.GLOG.Debugf(f, v...)
}
func Info(v ...interface{}) {
	base.GLOG.Info(v...)
}
func Infof(f string, v ...interface{}) {
	base.GLOG.Infof(f, v...)
}
func Warn(v ...interface{}) {
	base.GLOG.Warn(v...)
}
func Warnf(f string, v ...interface{}) {
	base.GLOG.Warnf(f, v...)
}

func Error(v ...interface{}) {
	base.GLOG.Error(v...)
}
func Errorf(f string, v ...interface{}) {
	base.GLOG.Errorf(f, v...)
}
func Fatal(v ...interface{}) {
	base.GLOG.Fatal(v...)
}
func Fatalf(f string, v ...interface{}) {
	base.GLOG.Fatalf(f, v...)
}
