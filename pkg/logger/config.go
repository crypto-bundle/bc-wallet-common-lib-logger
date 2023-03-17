package logger

// Config is config for logger service
type Config struct {
	// -------------------
	// Application configs
	// -------------------
	// MinimalLogsLevel is a level for setup minimal logger event notification.
	// Allowed values: debug, info, warn, error, dpanic, panic, fatal
	MinimalLogsLevel string `envconfig:"APP_LOGGER_LEVEL" default:"debug"`
	// StackTraceEnable is option for enable or disable stacktrace at zap.Logger level
	// Allowed values: true, false
	StackTraceEnable bool `envconfig:"APP_LOGGER_STACKTRACE_ENABLE" default:"false"`
	// -------------------
	// Graylog configs
	// -------------------
	GraylogEnable           bool   `envconfig:"GRAYLOG_ENABLE"`
	GraylogAddress          string `envconfig:"GRAYLOG_ADDRESS" default:"127.0.0.1:12201"`
	GraylogHost             string `envconfig:"GRAYLOG_HOST" default:"localhost"`
	GraylogVersion          string `envconfig:"GRAYLOG_VERSION" default:"1.1"`
	GraylogMinimalLogsLevel string `envconfig:"GRAYLOG_LOGGER_LEVEL" default:"warn"`
	// ---------------
	// Bugsnag configs
	// ---------------
	// BugsnagApiKey ...
	BugsnagEnable bool   `envconfig:"BUGSNAG_ENABLE"`
	BugsnagApiKey string `envconfig:"BUGSNAG_API_KEY" default:""`
	// BugsnagMinimalLogsLevel is a level for setup minimal bugsnag logger event notification
	// allowed: debug, info, warn, error, dpanic, panic, fatal
	BugsnagMinimalLogsLevel string `envconfig:"BUGSNAG_LOGGER_LEVEL" default:"warn"`
	// ---------------
	// Sentry configs
	// ---------------
	SentryEnable bool `envconfig:"SENTRY_ENABLE"`
	// SentryHost ...
	SentryHost string `envconfig:"SENTRY_HOST" default:"localhost"`
	// BugsnagMinimalLogsLevel is a level for setup minimal bugsnag logger event notification
	// allowed: debug, info, warn, error, dpanic, panic, fatal
	SentryMinimalLogsLevel string `envconfig:"SENTRY_LOGGER_LEVEL" default:"warn"`
}

func (c *Config) GetMinimalLogLevel() string {
	return c.MinimalLogsLevel
}

func (c *Config) IsStacktraceEnabled() bool {
	return c.StackTraceEnable
}
