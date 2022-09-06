// Copyright 2022 a-clap. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewDefaultZap(level zapcore.Level) *zap.SugaredLogger {
	cfg := zap.NewDevelopmentConfig()
	cfg.Level = zap.NewAtomicLevelAt(level)
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	log, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return log.Sugar()
}

func NewNop() *zap.SugaredLogger {
	return zap.NewNop().Sugar()
}

var _ Logger = NewDefaultZap(zapcore.DebugLevel)
var _ Logger = NewNop()
