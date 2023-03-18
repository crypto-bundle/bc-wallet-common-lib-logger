package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Service struct {
	cfg configManager

	defaultLogger *zap.Logger
	cores         []zapcore.Core

	entries map[string]*zap.Logger
}

func (s *Service) NewLoggerEntry(named string) (*zap.Logger, error) {
	var cores = []zapcore.Core{
		s.defaultLogger.Core(),
	}

	l := zap.New(zapcore.NewTee(cores...))
	zap.ReplaceGlobals(l)

	l = l.Named(named).With(zap.String(HostnameFieldTag, s.cfg.GetHostName()),
		zap.String(EnvironmentNameTag, s.cfg.GetEnvironmentName()),
		zap.String(StageNameTag, s.cfg.GetStageName()),
		zap.Int(ApplicationPID, s.cfg.GetApplicationPID()),
		zap.String(VersionTag, s.cfg.GetVersion()),
		zap.String(SVCReleaseTag, s.cfg.GetReleaseTag()),
		zap.String(SVCCommitShortID, s.cfg.GetShortCommitID()),
		zap.String(SVCCommitID, s.cfg.GetCommitID()),
		zap.Uint64(BuildNumberTag, s.cfg.GetBuildNumber()),
		zap.Time(BuildDateTag, s.cfg.GetBuildDate()),
		zap.Uint64(BuildDateTimestampTag, uint64(s.cfg.GetBuildDateTS())),
	)

	_, ok := s.entries[named]
	if ok {
		return nil, ErrNamedLoggerAlreadyRegistered
	}
	s.entries[named] = l

	return l, nil
}

func (s *Service) NewLoggerEntryWithFields(named string, fields ...zap.Field) (*zap.Logger, error) {
	var cores = []zapcore.Core{
		s.defaultLogger.Core(),
	}

	l := zap.New(zapcore.NewTee(cores...))
	zap.ReplaceGlobals(l)

	l = l.Named(named).With(fields...)

	_, ok := s.entries[named]
	if ok {
		return nil, ErrNamedLoggerAlreadyRegistered
	}
	s.entries[named] = l

	return l, nil
}

func NewService(cfg configManager) (*Service, error) {
	cores := make([]zapcore.Core, 1)

	logsLevel := new(zapcore.Level)
	err := logsLevel.Set(cfg.GetMinimalLogLevel())
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	lCfg := zap.NewProductionConfig()
	lCfg.Level = zap.NewAtomicLevelAt(*logsLevel)
	lCfg.DisableStacktrace = !cfg.IsStacktraceEnabled() // We use errs.ZapStack to get stacktrace
	lCfg.OutputPaths = []string{"stdout"}
	if cfg.IsDebug() {
		lCfg.Level.SetLevel(zap.DebugLevel)
	}

	defaultLogger, err := lCfg.Build()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	cores[0] = defaultLogger.Core()

	return &Service{
		cfg:           cfg,
		defaultLogger: defaultLogger,
		cores:         cores,
		entries:       make(map[string]*zap.Logger),
	}, nil
}
