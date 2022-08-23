// Copyright 2022 a-clap. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logger

type Logger interface {
	Errorf(format string, args ...interface{})
	Error(args ...interface{})

	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})

	Infof(format string, args ...interface{})
	Info(args ...interface{})

	Warnf(format string, args ...interface{})
	Warn(args ...interface{})

	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
}

var log Logger

func Init(logger Logger) {
	log = logger
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}
func Error(args ...interface{}) {
	log.Error(args...)
}
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}
func Info(args ...interface{}) {
	log.Info(args...)
}
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}
func Warn(args ...interface{}) {
	log.Warn(args...)
}
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}
func Debug(args ...interface{}) {
	log.Debug(args...)
}
