package logger

import (
	"log"
	"os"
	"runtime/debug"
	"time"

	"github.com/amh11706/qmail"
)

var infoColor = []interface{}{"\x1b[32m[INFO]\x1b[0m"}
var errorColor = []interface{}{"\x1b[31m[ERROR]\x1b[0m"}
var errorPrinter = log.New(os.Stderr, "", log.LstdFlags)

func Info(args ...interface{}) {
	log.Println(append(infoColor, args...)...)
}

func Error(args ...interface{}) {
	errorPrinter.Println(append(errorColor, args...)...)
}

var crashReportDebounce = make(chan struct{}, 1)

func CheckStack(err error) bool {
	failed := Check(err)
	if !failed {
		return failed
	}
	stack := string(debug.Stack())
	Error(stack)
	go func() {
		if len(crashReportDebounce) > 0 {
			return
		}
		Check(qmail.SendCrashReport(err, stack))
		<-crashReportDebounce
		time.Sleep(5 * time.Minute)
		crashReportDebounce <- struct{}{}
	}()
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
