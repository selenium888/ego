package elog

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestConfig_AtomicLevel(t *testing.T) {
	cfg := defaultConfig()
	cfg.al = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	assert.Equal(t, "info", cfg.al.String())
}

func TestConfig_EncoderConfig(t *testing.T) {
	cfg := defaultConfig()
	assert.Equal(t, "ts", cfg.encoderConfig.TimeKey)
	assert.Equal(t, "lv", cfg.encoderConfig.LevelKey)
	assert.Equal(t, "logger", cfg.encoderConfig.NameKey)
	assert.Equal(t, "caller", cfg.encoderConfig.CallerKey)
	assert.Equal(t, "msg", cfg.encoderConfig.MessageKey)
	assert.Equal(t, "stack", cfg.encoderConfig.StacktraceKey)
	assert.Equal(t, zapcore.DefaultLineEnding, cfg.encoderConfig.LineEnding)
}

func TestConfig_Filename(t *testing.T) {
	cfg := defaultConfig()
	assert.Equal(t, "./logs/default.log", cfg.Filename())
}

func Test_defaultConfig(t *testing.T) {
	cfg := defaultConfig()
	assert.Equal(t, DefaultLoggerName, cfg.Name)
	assert.Equal(t, "./logs", cfg.Dir)
	assert.Equal(t, "info", cfg.Level)
	assert.Equal(t, false, cfg.Debug)
	assert.Equal(t, false, cfg.EnableAddCaller)
	assert.Equal(t, true, cfg.EnableAsync)
	assert.Equal(t, "file", cfg.Writer)
	assert.Equal(t, 1, cfg.CallerSkip)
}
