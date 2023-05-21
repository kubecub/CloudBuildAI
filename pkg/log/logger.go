package log

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type LogOptions struct {
	// kubecub log file path, default log directory is `/var/lib/kubecub/log`
	OutputPath string
	// Verbose: kubecub log level,if it is ture will set debug log mode.
	Verbose bool
	// DisableColor if true will disable outputting colors.
	DisableColor         bool
	RemoteLoggerURL      string
	RemoteLoggerTaskName string
	// LogToFile flag represent whether write log to disk, default is false.
	LogToFile bool
}

func Init(options LogOptions) error {
	if options.Verbose {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.SetReportCaller(true)

	logrus.SetFormatter(&Formatter{
		DisableColor: options.DisableColor,
	})

	if options.LogToFile {
		fh, err := NewFileHook(options.OutputPath)
		if err != nil {
			return errors.Errorf("failed to init log file hook: %v", err)
		}
		logrus.AddHook(fh)
	}

	if options.RemoteLoggerURL != "" {
		rl, err := NewRemoteLogHook(options.RemoteLoggerURL, options.RemoteLoggerTaskName)
		if err != nil {
			return errors.Errorf("failed to init log remote hook: %v", err)
		}
		logrus.AddHook(rl)
	}

	return nil
}
