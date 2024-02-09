package logger

// LoggerConfig is config for logger service
type LoggerConfig struct {
	// -------------------
	// Application configs
	// -------------------
	// MinimalLogsLevel is a level for setup minimal logger event notification.
	// Allowed values: debug, info, warn, error, dpanic, panic, fatal
	MinimalLogsLevel string `envconfig:"LOGGER_LEVEL" default:"debug"`
	// StackTraceEnable is option for enable or disable stacktrace at zap.Logger level
	// Allowed values: true, false
	StackTraceEnable bool `envconfig:"LOGGER_STACKTRACE_ENABLE" default:"false"`
}

func (c *LoggerConfig) GetMinimalLogLevel() string {
	return c.MinimalLogsLevel
}

func (c *LoggerConfig) IsStacktraceEnabled() bool {
	return c.StackTraceEnable
}
