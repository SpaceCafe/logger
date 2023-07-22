package logger

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger()

	if logger.level != LevelDebug {
		t.Error("expected logger level to be debug")
	}

	if logger.loggerList[LevelDebug] == nil {
		t.Error("expected logger loggerList[0] to be initialized")
	}

	if logger.loggerList[LevelInfo] == nil {
		t.Error("expected logger loggerList[1] to be initialized")
	}

	if logger.loggerList[LevelWarn] == nil {
		t.Error("expected logger loggerList[2] to be initialized")
	}

	if logger.loggerList[LevelFatal] == nil {
		t.Error("expected logger loggerList[3] to be initialized")
	}
}

func TestLogger_Level(t *testing.T) {
	tests := []LogLevel{LevelDebug, LevelInfo, LevelWarn, LevelFatal}

	for _, tt := range tests {
		logger := NewLogger()
		logger.SetLevel(tt)

		assert.Equal(t, tt, logger.Level())
	}
}

func FuzzLogger_Debug(f *testing.F) {
	tests := []string{"test", "test2\n", "test 👍"}
	for _, tt := range tests {
		f.Add(tt)
	}

	f.Fuzz(func(t *testing.T, tt string) {
		var buf bytes.Buffer
		logger := NewLogger()
		logger.loggerList[LevelDebug].SetOutput(&buf)
		logger.loggerList[LevelDebug].SetFlags(0)

		logger.Debug(tt)

		if len(tt) > 0 && tt[len(tt)-1] == '\n' {
			assert.Equalf(t, "[\u001B[0;37mDEBUG\u001B[0m]   "+tt, buf.String(), "expected logger to write %#v to stdout", tt)
		} else {
			assert.Equalf(t, "[\u001B[0;37mDEBUG\u001B[0m]   "+tt+"\n", buf.String(), "expected logger to write %#v to stdout", tt)
		}
	})
}

func TestLogger_Print(t *testing.T) {
	tests := []struct {
		name  string
		level LogLevel
		args  []any
	}{
		{"debug", LevelDebug, []any{"test", 2, "test 👍"}},
		{"info", LevelInfo, []any{"test", 2, "test 👍"}},
		{"warn", LevelWarn, []any{"test", 2, "test 👍"}},
		{"fatal", LevelFatal, []any{"test", 2, "test 👍"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var buf0 bytes.Buffer
			var buf1 bytes.Buffer
			var buf2 bytes.Buffer
			var buf3 bytes.Buffer

			logger := NewLogger()
			osExit = func(code int) {}
			logger.loggerList[LevelDebug].SetOutput(&buf0)
			logger.loggerList[LevelDebug].SetFlags(0)
			logger.loggerList[LevelInfo].SetOutput(&buf1)
			logger.loggerList[LevelInfo].SetFlags(0)
			logger.loggerList[LevelWarn].SetOutput(&buf2)
			logger.loggerList[LevelWarn].SetFlags(0)
			logger.loggerList[LevelFatal].SetOutput(&buf3)
			logger.loggerList[LevelFatal].SetFlags(0)

			logger.SetLevel(tt.level)

			logger.Debug(tt.args...)
			logger.Info(tt.args...)
			logger.Warn(tt.args...)
			logger.Fatal(tt.args...)

			logger.Debugf("%v", tt.args...)
			logger.Infof("%v", tt.args...)
			logger.Warnf("%v", tt.args...)
			logger.Fatalf("%v", tt.args...)

			if tt.level <= LevelDebug {
				assert.Equalf(t, "[\u001B[0;37mDEBUG\u001B[0m]   "+fmt.Sprint(tt.args...)+"\n[\u001B[0;37mDEBUG\u001B[0m]   "+fmt.Sprintf("%v", tt.args...)+"\n", buf0.String(), "expected logger to write %#v to stdout", tt.args)
			} else {
				assert.Equalf(t, "", buf0.String(), "expected logger to not write %#v to stdout", tt.args)
			}

			if tt.level <= LevelInfo {
				assert.Equalf(t, "[\u001B[0;32mINFO\u001B[0m]    "+fmt.Sprint(tt.args...)+"\n[\u001B[0;32mINFO\u001B[0m]    "+fmt.Sprintf("%v", tt.args...)+"\n", buf1.String(), "expected logger to write %#v to stdout", tt.args)
			} else {
				assert.Equalf(t, "", buf1.String(), "expected logger to not write %#v to stdout", tt.args)
			}

			if tt.level <= LevelWarn {
				assert.Equalf(t, "[\u001B[0;33mWARNING\u001B[0m] "+fmt.Sprint(tt.args...)+"\n[\u001B[0;33mWARNING\u001B[0m] "+fmt.Sprintf("%v", tt.args...)+"\n", buf2.String(), "expected logger to write %#v to stdout", tt.args)
			} else {
				assert.Equalf(t, "", buf2.String(), "expected logger to not write %#v to stdout", tt.args)
			}

			if tt.level <= LevelFatal {
				assert.Equalf(t, "[\u001B[0;31mFATAL\u001B[0m]   "+fmt.Sprint(tt.args...)+"\n[\u001B[0;31mFATAL\u001B[0m]   "+fmt.Sprintf("%v", tt.args...)+"\n", buf3.String(), "expected logger to write %#v to stdout", tt.args)
			} else {
				assert.Equalf(t, "", buf3.String(), "expected logger to not write %#v to stdout", tt.args)
			}
		})
	}
}

type mockBuf []byte

func (b *mockBuf) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("test error")
}

func TestLogger_logWithLevelErr(t *testing.T) {
	buf := make(mockBuf, 0)
	logger := NewLogger()
	logger.loggerList[LevelDebug].SetOutput(&buf)
	logger.loggerList[LevelDebug].SetFlags(0)

	logger.logWithLevel(LevelDebug, nil, "test")
	assert.Equal(t, 0, len(buf), "expected logger to not write to stdout")
}
