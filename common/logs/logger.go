package logs

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

var zapOnce sync.Once
var innerSugarLogger *zap.SugaredLogger

func InitNomadGlobalLog() {
	zapOnce.Do(func() {
		pwdDir, _ := os.Getwd()
		infoLogFilePath := pwdDir + "/alloc/logs/demo-log-main-task.stdout.0"
		errLogFilePath := pwdDir + "/alloc/logs/demo-log-main-task.stderr.0"
		serverName := "demo-log"
		if v := os.Getenv("APP_NAME"); v != "" {
			serverName = v
		}
		branchName := "main"
		if v := os.Getenv("CI_COMMIT_BRANCH"); v != "" {
			branchName = v
		} else {
			branchName = "main"
		}
		infoLogFilePath = strings.ReplaceAll(infoLogFilePath, "/app", "")
		errLogFilePath = strings.ReplaceAll(errLogFilePath, "/app", "")
		// 替换 serverName
		infoLogFilePath = strings.ReplaceAll(infoLogFilePath, "demo-log", serverName)
		errLogFilePath = strings.ReplaceAll(errLogFilePath, "demo-log", serverName)
		// 替换 branchName
		infoLogFilePath = strings.ReplaceAll(infoLogFilePath, "main", branchName)
		errLogFilePath = strings.ReplaceAll(errLogFilePath, "main", branchName)

		// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
		encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			MessageKey:  "msg",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			TimeKey:     "ts",
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02 15:04:05"))
			},
			EncodeName:   zapcore.FullNameEncoder,
			CallerKey:    "file",
			EncodeCaller: zapcore.ShortCallerEncoder,
			EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendInt64(int64(d) / 1000000)
			},
		})

		// 实现两个判断日志等级的interface
		infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.InfoLevel
		})

		errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.ErrorLevel
		})

		// 获取 info、error日志文件的io.Writer 抽象 getWriter() 在下方实现
		infoWriter := getWriter(infoLogFilePath)
		errorWriter := getWriter(errLogFilePath)

		// 最后创建具体的Logger
		core := zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
		)

		logger := zap.New(core, zap.AddCaller()) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑
		innerSugarLogger = logger.Sugar().With("appName", serverName)
	})
}

type logger struct {
	log     *zap.SugaredLogger
	traceId string
}

func NewLog(traceId string) *logger {
	if innerSugarLogger == nil {
		InitNomadGlobalLog()
	}

	if traceId == "" {
		return &logger{
			log: innerSugarLogger,
		}
	}

	return &logger{
		log:     innerSugarLogger.With(zap.String("trace_id", traceId)),
		traceId: traceId,
	}
}

func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		strings.Replace(filename, ".0", "", -1)+".%Y%m%d%H", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24),
		rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}
	return hook
}

func (l logger) Debug(args ...interface{}) {
	l.log.Debug(args...)
}

func (l logger) Debugf(template string, args ...interface{}) {
	l.log.Debugf(template, args...)
}

func (l logger) Info(args ...interface{}) {
	l.log.Info(args...)
}

func (l logger) Infof(template string, args ...interface{}) {
	l.log.Infof(template, args...)
}

func (l logger) Warn(args ...interface{}) {
	l.log.Warn(args...)
}

func (l logger) Warnf(template string, args ...interface{}) {
	l.log.Warnf(template, args...)
}

func (l logger) Error(args ...interface{}) {
	l.log.Error(args...)
}

func (l logger) Errorf(template string, args ...interface{}) {
	l.log.Errorf(template, args...)
}

func (l logger) DPanic(args ...interface{}) {
	l.log.DPanic(args...)
}

func (l logger) DPanicf(template string, args ...interface{}) {
	l.log.DPanicf(template, args...)
}

func (l logger) Panic(args ...interface{}) {
	l.log.Panic(args...)
}

func (l logger) Panicf(template string, args ...interface{}) {
	l.log.Panicf(template, args...)
}

func (l logger) Fatal(args ...interface{}) {
	l.log.Fatal(args...)
}

func (l logger) Fatalf(template string, args ...interface{}) {
	l.log.Fatalf(template, args...)
}
