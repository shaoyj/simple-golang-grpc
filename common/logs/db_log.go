package logs

import (
	c "context"
	"go.uber.org/zap"
	"gorm.io/gorm"
	llog "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

const LOG_NAMESPACE = "_gorm"

func NewZapLogger() llog.Interface {
	zapLogger, _ := zap.NewProduction()
	return &ZapLogger{
		defaultLogger: zapLogger,
		level:         llog.Info,
	}
}

type ZapLogger struct {
	level                     llog.LogLevel
	defaultLogger             *zap.Logger
	IgnoreRecordNotFoundError bool
}

func (z *ZapLogger) LogMode(level llog.LogLevel) llog.Interface {
	newlogger := *z
	newlogger.level = level
	return &newlogger
}

func (l ZapLogger) Info(ctx c.Context, msg string, data ...interface{}) {
	ll := l.defaultLogger

	ll.Info(msg, zap.Namespace(LOG_NAMESPACE), zap.String("caller", utils.FileWithLineNum()), zap.String("msg", msg), zap.Any("data", data))
}

func (l ZapLogger) Warn(ctx c.Context, msg string, data ...interface{}) {
	ll := l.defaultLogger

	ll.Warn(msg, zap.Namespace(LOG_NAMESPACE), zap.String("caller", utils.FileWithLineNum()), zap.String("msg", msg), zap.Any("data", data))
}

func (l ZapLogger) Error(ctx c.Context, msg string, data ...interface{}) {
	ll := l.defaultLogger

	ll.Error(msg, zap.Namespace(LOG_NAMESPACE), zap.String("caller", utils.FileWithLineNum()), zap.String("msg", msg), zap.Any("data", data))
}

func (l ZapLogger) Trace(ctx c.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	ll := l.defaultLogger

	if nil == err {
		if l.level < llog.Info {
			return
		}

		ll.With(zap.Int64("cost", elapsed.Milliseconds()), zap.String("sql", sql)).
			Info("gorm trace", zap.Namespace(LOG_NAMESPACE),
				zap.String("caller", utils.FileWithLineNum()),
				zap.Int64("rows", rows))
	} else {
		//数据不存在 不需要报错
		if err == gorm.ErrRecordNotFound {
			return
		}

		ll.With(zap.Int64("cost", elapsed.Milliseconds()), zap.String("sql", sql)).
			Error(err.Error(), zap.Namespace(LOG_NAMESPACE),
				zap.String("caller", utils.FileWithLineNum()),
				zap.Error(err),
				zap.Int64("rows", rows))
	}
}
