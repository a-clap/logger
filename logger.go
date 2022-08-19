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
