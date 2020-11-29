package logger

import (
	"log"
	"runtime/debug"
)

var infoColor = []interface{}{"\x1b[32m[INFO]\x1b[0m"}
var errorColor = []interface{}{"\x1b[31m[ERROR]\x1b[0m"}

func Info(args ...interface{}) {
	log.Println(append(infoColor, args...)...)
}

func Error(args ...interface{}) {
	log.Println(append(errorColor, args...)...)
}

func CheckStack(err error) bool {
	failed := Check(err)
	if failed {
		debug.PrintStack()
	}
	return failed
}

func Check(err error) bool {
	if err != nil {
		Error(err)
		return true
	}
	return false
}

func CheckP(err error, prefix string) bool {
	if err != nil {
		Error(prefix, err)
		return true
	}
	return false
}
