package logger

import "time"

type configManager interface {
	GetHostName() string
	GetEnvironmentName() string
	IsProd() bool
	IsStage() bool
	IsTest() bool
	IsDev() bool
	IsDebug() bool
	IsLocal() bool
	GetStageName() string

	GetApplicationPID() int
	GetReleaseTag() string
	GetCommitID() string
	GetShortCommitID() string
	GetBuildNumber() uint64
	GetBuildDateTS() int64
	GetBuildDate() time.Time

	GetMinimalLogLevel() string
	GetSkipBuildInfo() bool
	IsStacktraceEnabled() bool
}
