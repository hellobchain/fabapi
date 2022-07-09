package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/signal"
	"runtime/pprof"
	"strings"
	"syscall"

	"github.com/wsw365904/fabapi/core/common/e"

	"github.com/pkg/errors"
)

// 第三方包和自定义错误，封装stack
func ToErr(err error, format string, args ...interface{}) error {
	if err == nil {
		return errors.Errorf(format, args...)
	}
	return errors.Wrapf(err, format, args...)
}

// 打印详细信息，附带堆栈信息。
func PrintStack(err error) string {
	errMsg := fmt.Sprintf("%+v", err)
	return cleanPath(errMsg)
}

// 脱敏
func cleanPath(s string) string {
	return strings.ReplaceAll(s, getCurrentPath()+"/", "")
}

// 获取当前项目目录
func getCurrentPath() string {
	getwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return strings.Replace(getwd, "\\", "/", -1)
}

// 给第三方包添加本地错误编码
// 成功或者nil直接返回不需要包装
func NewError(code e.ErrCode, err error) error {
	if code == e.SUCCESS {
		return nil
	}
	if err != nil {
		return errors.Wrapf(code, err.Error())
	}
	return code
}

func HandleSignals(handlers map[os.Signal]func(), logger Logger) {
	var signals []os.Signal
	for sig := range handlers {
		signals = append(signals, sig)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, signals...)

	for sig := range signalChan {
		logger.Infof("Received signal: %d (%s)", sig, sig)
		handlers[sig]()
	}
}

func AddPlatformSignals(sigs map[os.Signal]func(), logger Logger) map[os.Signal]func() {
	sigs[syscall.SIGUSR1] = func() { logGoRoutines(logger) }
	return sigs
}

type Logger interface {
	Infof(template string, args ...interface{})
	Errorf(template string, args ...interface{})
}

func captureGoRoutines() (string, error) {
	var buf bytes.Buffer
	err := pprof.Lookup("goroutine").WriteTo(&buf, 2)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func logGoRoutines(logger Logger) {
	output, err := captureGoRoutines()
	if err != nil {
		logger.Errorf("failed to capture go routines: %s", err)
		return
	}

	logger.Infof("Go routines report:\n%s", output)
}
