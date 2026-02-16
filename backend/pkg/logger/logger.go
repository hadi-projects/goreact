package logger

import (
	"io"
	"os"
	"path/filepath"

	"github.com/hadi-projects/go-react-starter/config"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger interface {
	Info() *zerolog.Event
	Error() *zerolog.Event
	Debug() *zerolog.Event
	Warn() *zerolog.Event
	Fatal() *zerolog.Event
}

type logger struct {
	zLog zerolog.Logger
}

func (l *logger) Info() *zerolog.Event {
	return l.zLog.Info()
}

func (l *logger) Error() *zerolog.Event {
	return l.zLog.Error()
}

func (l *logger) Debug() *zerolog.Event {
	return l.zLog.Debug()
}

func (l *logger) Warn() *zerolog.Event {
	return l.zLog.Warn()
}

func (l *logger) Fatal() *zerolog.Event {
	return l.zLog.Fatal()
}

// Global instances for backward compatibility if needed,
// strictly speaking we should move to DI for everything.
var (
	SystemLogger    zerolog.Logger
	AuthLogger      zerolog.Logger
	DBLogger        zerolog.Logger
	RedisLogger     zerolog.Logger
	RateLimitLogger zerolog.Logger
)

func InitLogger(cfg *config.Config) {
	if err := os.MkdirAll(cfg.Log.Dir, 0755); err != nil {
		panic(err)
	}

	SystemLogger = newZeroLogger(*cfg, "system.log")
	AuthLogger = newZeroLogger(*cfg, "auth.log")
	DBLogger = newZeroLogger(*cfg, "db.log")
	RedisLogger = newZeroLogger(*cfg, "redis.log")
	RateLimitLogger = newZeroLogger(*cfg, "rate_limit.log")
}

func NewLogger(cfg config.Config, fileName string) Logger {
	zLog := newZeroLogger(cfg, fileName)
	return &logger{zLog: zLog}
}

func newZeroLogger(cfg config.Config, fileName string) zerolog.Logger {
	fileLogger := &lumberjack.Logger{
		Filename:   filepath.Join(cfg.Log.Dir, fileName),
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}

	var writers []io.Writer
	writers = append(writers, fileLogger)

	if cfg.App.Env == "development" {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stdout})
	}

	multi := io.MultiWriter(writers...)

	return zerolog.New(multi).With().Timestamp().Logger()
}
