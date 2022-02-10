package global

import (
	"github.com/go-redis/redis/v8"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/spf13/cobra"
	"go-scaffold/internal/app/config"
	oteltrace "go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Global 是传递给应用的依赖
var (
	command      *cobra.Command         // 根命令实例
	loggerOutput *rotatelogs.RotateLogs // 日志输出实例
	conf         *config.Config         // 配置实例
	logger       *zap.Logger            // 日志实例
	db           *gorm.DB               // 数据库实例
	redisClient  *redis.Client          // redis 实例
	tracer       oteltrace.Tracer       // Tracer 实例
)

// SetCommand 设置命令行实例
func SetCommand(cmd *cobra.Command) {
	command = cmd
}

// Command 获取命令行实例
func Command() *cobra.Command {
	return command
}

// SetLoggerOutput 设置日志轮转实例
func SetLoggerOutput(lr *rotatelogs.RotateLogs) {
	loggerOutput = lr
}

// LoggerOutput 获取日志轮转实例
func LoggerOutput() *rotatelogs.RotateLogs {
	return loggerOutput
}

// SetConfig 设置日志实例
func SetConfig(c *config.Config) {
	conf = c
}

// Config 获取日志实例
func Config() *config.Config {
	return conf
}

// SetLogger 设置日志实例
func SetLogger(l *zap.Logger) {
	logger = l
}

// Logger 获取日志实例
func Logger() *zap.Logger {
	return logger
}

// SetDB 设置数据库实例
func SetDB(d *gorm.DB) {
	db = d
}

// DB 获取数据库实例
func DB() *gorm.DB {
	return db
}

// SetRedisClient 设置 redis 客户端实例
func SetRedisClient(rc *redis.Client) {
	redisClient = rc
}

// RedisClient 获取 redis 客户端实例
func RedisClient() *redis.Client {
	return redisClient
}

// SetTracer 设置 Tracer 实例
func SetTracer(t oteltrace.Tracer) {
	tracer = t
}

// Tracer 获取 Tracer 实例
func Tracer() oteltrace.Tracer {
	return tracer
}
