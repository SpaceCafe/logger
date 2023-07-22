package logger

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestDefault(t *testing.T) {
	logger := NewLogger()
	assert.Equal(t, reflect.DeepEqual(Default(), logger), true, "expected Default() to be the same as NewLogger()")
	assert.Equal(t, reflect.ValueOf(Level).Pointer(), reflect.ValueOf(logger.Level).Pointer(), "expected Level() to be the same as logger.Level()")
	assert.Equal(t, reflect.ValueOf(SetLevel).Pointer(), reflect.ValueOf(logger.SetLevel).Pointer(), "expected SetLevel() to be the same as logger.SetLevel()")
	assert.Equal(t, reflect.ValueOf(Debug).Pointer(), reflect.ValueOf(logger.Debug).Pointer(), "expected Debug() to be the same as logger.Debug()")
	assert.Equal(t, reflect.ValueOf(Debugf).Pointer(), reflect.ValueOf(logger.Debugf).Pointer(), "expected Debugf() to be the same as logger.Debugf()")
	assert.Equal(t, reflect.ValueOf(Info).Pointer(), reflect.ValueOf(logger.Info).Pointer(), "expected Info() to be the same as logger.Info()")
	assert.Equal(t, reflect.ValueOf(Infof).Pointer(), reflect.ValueOf(logger.Infof).Pointer(), "expected Infof() to be the same as logger.Infof()")
	assert.Equal(t, reflect.ValueOf(Warn).Pointer(), reflect.ValueOf(logger.Warn).Pointer(), "expected Warn() to be the same as logger.Warn()")
	assert.Equal(t, reflect.ValueOf(Warnf).Pointer(), reflect.ValueOf(logger.Warnf).Pointer(), "expected Warnf() to be the same as logger.Warnf()")
	assert.Equal(t, reflect.ValueOf(Fatal).Pointer(), reflect.ValueOf(logger.Fatal).Pointer(), "expected Fatal() to be the same as logger.Fatal()")
	assert.Equal(t, reflect.ValueOf(Fatalf).Pointer(), reflect.ValueOf(logger.Fatalf).Pointer(), "expected Fatalf() to be the same as logger.Fatalf()")
}
