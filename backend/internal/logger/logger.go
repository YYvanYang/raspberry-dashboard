package logger

import (
    "fmt"
    "os"
    "path/filepath"
    "time"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
)

var log *zap.SugaredLogger

func Init(logPath string) error {
    // 确保日志目录存在
    if err := os.MkdirAll(filepath.Dir(logPath), 0750); err != nil {
        return fmt.Errorf("failed to create log directory: %w", err)
    }

    // 配置日志轮转
    writer := &lumberjack.Logger{
        Filename:   logPath,
        MaxSize:    10,    // 每个日志文件最大尺寸，单位是 MB
        MaxBackups: 5,     // 保留的旧日志文件数量
        MaxAge:     30,    // 保留的最大天数
        Compress:   true,  // 压缩旧日志文件
    }

    // 日志编码配置
    encoderConfig := zapcore.EncoderConfig{
        TimeKey:        "time",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "caller",
        MessageKey:     "msg",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.LowercaseLevelEncoder,
        EncodeTime:     zapcore.ISO8601TimeEncoder,
        EncodeDuration: zapcore.SecondsDurationEncoder,
        EncodeCaller:   zapcore.ShortCallerEncoder,
    }

    // 创建核心
    core := zapcore.NewCore(
        zapcore.NewJSONEncoder(encoderConfig),
        zapcore.AddSync(writer),
        zap.NewAtomicLevelAt(zap.InfoLevel),
    )

    // 创建记录器
    logger := zap.New(core, zap.AddCaller())
    log = logger.Sugar()
    return nil
}

// 日志方法
func Info(msg string, args ...interface{})  { log.Infof(msg, args...) }
func Error(msg string, args ...interface{}) { log.Errorf(msg, args...) }
func Warn(msg string, args ...interface{})  { log.Warnf(msg, args...) }
func Debug(msg string, args ...interface{}) { log.Debugf(msg, args...) } 