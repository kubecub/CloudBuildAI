package logger

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	"github.com/kubecub/CloudBuildAI/pkg/common"
)

func NewFileHook(filePath string) (logrus.Hook, error) {
	if filePath == "" {
		filePath = common.DefaultLogDir
	}

	path := filepath.Join(filePath, "kubecub.log")
	writer, _ := rotatelogs.New(
		path+".%Y%m%d",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	return lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.DebugLevel: writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{
		DisableColors: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return "", fmt.Sprintf("%s:%d", filepath.Base(frame.File), frame.Line)
		},
	}), nil
}
