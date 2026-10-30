package log

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func isolatedLogger(t *testing.T) *logrus.Logger {
	t.Helper()
	logger := logrus.StandardLogger()
	output, formatter, level, caller, hooks := logger.Out, logger.Formatter, logger.GetLevel(), logger.ReportCaller, logger.Hooks
	logger.ReplaceHooks(make(logrus.LevelHooks))
	t.Cleanup(func() {
		logger.SetOutput(output)
		logger.SetFormatter(formatter)
		logger.SetLevel(level)
		logger.SetReportCaller(caller)
		logger.ReplaceHooks(hooks)
	})
	return logger
}

func TestInitConfiguresActualLogrusLogger(t *testing.T) {
	logger := isolatedLogger(t)
	var output bytes.Buffer
	logger.SetOutput(&output)
	if err := Init(LogOptions{Verbose: true, DisableColor: true}); err != nil {
		t.Fatal(err)
	}
	if logger.GetLevel() != logrus.DebugLevel || !logger.ReportCaller {
		t.Fatal("Init did not configure level and caller reporting")
	}
	formatter, ok := logger.Formatter.(*Formatter)
	if !ok || !formatter.DisableColor {
		t.Fatal("Init did not install the project formatter")
	}
	Errorf("synthetic %s", "message")
	if !strings.Contains(output.String(), "[ERROR]") || !strings.Contains(output.String(), "synthetic message") {
		t.Fatalf("Errorf did not use the configured logger: %q", output.String())
	}
}

type baselineHook struct {
	calls int
}

func (h *baselineHook) Levels() []logrus.Level { return []logrus.Level{logrus.WarnLevel} }

func (h *baselineHook) Fire(*logrus.Entry) error {
	h.calls++
	return nil
}

func TestInitRejectsRemoteLoggingWithoutChangingLocalLogger(t *testing.T) {
	logger := isolatedLogger(t)
	var output bytes.Buffer
	logger.SetOutput(&output)
	if err := Init(LogOptions{DisableColor: true}); err != nil {
		t.Fatal(err)
	}
	hook := &baselineHook{}
	logger.AddHook(hook)
	level, formatter, caller := logger.GetLevel(), logger.Formatter, logger.ReportCaller
	hooks := logger.Hooks
	snapshot := make(logrus.LevelHooks, len(hooks))
	for level, entries := range hooks {
		snapshot[level] = append([]logrus.Hook(nil), entries...)
	}
	path := filepath.Join(t.TempDir(), "not-created", "logs")
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("expected absent log directory before Init, got %v", err)
	}
	err := Init(LogOptions{
		RemoteLoggerURL: "https://synthetic.invalid/logger",
		LogToFile:       true,
		OutputPath:      path,
		Verbose:         true,
	})
	if err == nil || err.Error() != "remote logging is not supported" {
		t.Fatalf("unsupported remote logging must fail explicitly: %v", err)
	}
	if logger.GetLevel() != level || logger.Formatter != formatter || logger.ReportCaller != caller || logger.Out != &output {
		t.Fatal("rejected remote logging changed the established global logger")
	}
	if len(logger.Hooks) != len(snapshot) {
		t.Fatal("rejected remote logging changed the established hook map")
	}
	for level, entries := range snapshot {
		got := logger.Hooks[level]
		if len(got) != len(entries) {
			t.Fatalf("rejected remote logging changed hooks for level %s", level)
		}
		for i := range entries {
			if got[i] != entries[i] {
				t.Fatal("rejected remote logging replaced an established hook")
			}
		}
	}
	// Mutating the saved map checks map identity without relying on private fields.
	probe := &baselineHook{}
	hooks[logrus.TraceLevel] = []logrus.Hook{probe}
	if got := logger.Hooks[logrus.TraceLevel]; len(got) != 1 || got[0] != probe {
		t.Fatal("rejected remote logging replaced the established hook map")
	}
	delete(hooks, logrus.TraceLevel)
	logger.Warn("synthetic baseline logger")
	if hook.calls != 1 || !strings.Contains(output.String(), "synthetic baseline logger") {
		t.Fatal("the established local logger no longer works after rejection")
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("rejected remote logging created a log directory: %v", err)
	}
}

func TestFormatterPreservesCallerAndText(t *testing.T) {
	logger := logrus.New()
	logger.SetReportCaller(true)
	entry := logrus.NewEntry(logger)
	entry.Message, entry.Level = "synthetic message", logrus.WarnLevel
	entry.Caller = &runtime.Frame{File: "/synthetic/path/caller.go", Line: 42, Function: "synthetic.caller"}
	got, err := (&Formatter{DisableColor: true, HideLogTime: true}).Format(entry)
	if err != nil || string(got) != " [WARNING] [caller.go:42] synthetic message\n\n" {
		t.Fatalf("formatter changed caller/text bytes: %q, %v", got, err)
	}
}

type writerOutput chan string

func (w writerOutput) Write(p []byte) (int, error) {
	w <- string(p)
	return len(p), nil
}

func TestWriterContract(t *testing.T) {
	logger := isolatedLogger(t)
	if err := Init(LogOptions{DisableColor: true}); err != nil {
		t.Fatal(err)
	}
	logger.SetFormatter(&Formatter{DisableColor: true, HideLogTime: true, HideLogPath: true})
	for _, test := range []struct {
		name, input string
		want        []string
	}{
		{name: "lines", input: "alpha\nbeta\n", want: []string{"alpha", "beta"}},
		{name: "CRLF and partial EOF", input: "alpha\r\nbeta\r\ntail", want: []string{"alpha", "beta", "tail"}},
		{name: "above scanner limit", input: strings.Repeat("x", 2*65536+7)},
	} {
		t.Run(test.name, func(t *testing.T) {
			output := make(writerOutput, 16)
			logger.SetOutput(output)
			writer := logger.WriterLevel(logrus.ErrorLevel)
			t.Cleanup(func() { _ = writer.Close() })
			done := make(chan error, 1)
			go func() {
				_, err := io.WriteString(writer, test.input)
				closeErr := writer.Close()
				if err == nil {
					err = closeErr
				}
				done <- err
			}()
			select {
			case err := <-done:
				if err != nil {
					t.Fatalf("Writer write/close failed: %v", err)
				}
			case <-time.After(3 * time.Second):
				t.Fatal("Writer write/close timed out")
			}
			var got []string
			deadline := time.After(3 * time.Second)
			for {
				select {
				case text := <-output:
					if !strings.HasPrefix(text, " [ERROR] ") || !strings.HasSuffix(text, "\n\n") {
						t.Fatalf("Writer lost formatter/level: %q", text)
					}
					got = append(got, strings.TrimSuffix(strings.TrimPrefix(text, " [ERROR] "), "\n\n"))
				case <-deadline:
					t.Fatalf("Writer did not deliver complete content: %d entries", len(got))
				}
				if test.want != nil && len(got) == len(test.want) {
					break
				}
				if test.want == nil && len(strings.Join(got, "")) >= len(test.input) {
					break
				}
			}
			if test.want != nil {
				if strings.Join(got, "|") != strings.Join(test.want, "|") {
					t.Fatalf("Writer changed line/EOF contract: %q", got)
				}
			} else if strings.Join(got, "") != test.input {
				t.Fatal("Writer lost, duplicated or changed long input")
			}
		})
	}
}
