package utils

import (
	"os"

	"go.uber.org/zap"
)

var Logger *zap.Logger

var Sugar *zap.SugaredLogger

func init() {
	Logger, _ = zap.NewDevelopment()
	Sugar = Logger.Sugar()
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	Sugar.Error(err)
	os.Exit(1)
}

// Infof should be used to describe the example commands that are about to run.
func Infof(format string, args ...interface{}) {
	Sugar.Infof(format, args...)
}

// Warningf should be used to display a warning.
func Warningf(format string, args ...interface{}) {
	Sugar.Warnf(format, args...)
}

// Errorf should be used to display an error.
func Errorf(format string, args ...interface{}) {
	Sugar.Errorf(format, args...)
}

// Debugf should be used to output debug information.
func Debugf(format string, args ...interface{}) {
	Sugar.Debugf(format, args...)
}
