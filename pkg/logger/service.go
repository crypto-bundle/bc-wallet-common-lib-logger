/*
 *
 *
 * MIT NON-AI License
 *
 * Copyright (c) 2022-2024 Aleksei Kotelnikov(gudron2s@gmail.com)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of the software and associated documentation files (the "Software"),
 * to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions.
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * In addition, the following restrictions apply:
 *
 * 1. The Software and any modifications made to it may not be used for the purpose of training or improving machine learning algorithms,
 * including but not limited to artificial intelligence, natural language processing, or data mining. This condition applies to any derivatives,
 * modifications, or updates based on the Software code. Any usage of the Software in an AI-training dataset is considered a breach of this License.
 *
 * 2. The Software may not be included in any dataset used for training or improving machine learning algorithms,
 * including but not limited to artificial intelligence, natural language processing, or data mining.
 *
 * 3. Any person or organization found to be in violation of these restrictions will be subject to legal action and may be held liable
 * for any damages resulting from such use.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
 * DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
 * OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 *
 */

package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Service struct {
	cfg configManager

	defaultLogger *zap.Logger
	cores         []zapcore.Core
}

func (s *Service) NewLoggerEntry(named string) *zap.Logger {
	return s.newLoggerEntry(named)
}

func (s *Service) newLoggerEntry(named string) *zap.Logger {
	var cores = []zapcore.Core{
		s.defaultLogger.Core(),
	}

	l := zap.New(zapcore.NewTee(cores...))
	zap.ReplaceGlobals(l)

	l = l.Named(named).With(zap.String(HostnameFieldTag, s.cfg.GetHostName()),
		zap.String(EnvironmentNameTag, s.cfg.GetEnvironmentName()),
		zap.String(StageNameTag, s.cfg.GetStageName()),
		zap.Int(ApplicationPID, s.cfg.GetApplicationPID()))

	isDevOrLocal := s.cfg.IsDev() || s.cfg.IsLocal()
	buildInfoEnabled := isDevOrLocal && s.cfg.GetSkipBuildInfo()
	if buildInfoEnabled {
		l = l.With(zap.String(SVCReleaseTag, s.cfg.GetReleaseTag()),
			zap.String(SVCCommitShortID, s.cfg.GetShortCommitID()),
			zap.String(SVCCommitID, s.cfg.GetCommitID()),
			zap.Uint64(BuildNumberTag, s.cfg.GetBuildNumber()),
			zap.Time(BuildDateTag, s.cfg.GetBuildDate()),
			zap.Uint64(BuildDateTimestampTag, uint64(s.cfg.GetBuildDateTS())))
	}

	return l
}

func (s *Service) NewLoggerEntryWithFields(named string, fields ...zap.Field) *zap.Logger {
	var cores = []zapcore.Core{
		s.defaultLogger.Core(),
	}

	l := zap.New(zapcore.NewTee(cores...))
	zap.ReplaceGlobals(l)

	l = l.Named(named).With(fields...)

	return l
}

func (s *Service) NewStdLogMaker() *stdLogFabric {
	return &stdLogFabric{
		zapMakerFunc: s.newLoggerEntry,
	}
}

func (s *Service) NewNamedStdLogMaker(name string) *singleNameStdLogFabric {
	return &singleNameStdLogFabric{
		zapMakerFunc: s.newLoggerEntry,
		name:         name,
	}
}

func NewService(cfg configManager) (*Service, error) {
	cores := make([]zapcore.Core, 1)

	logsLevel := new(zapcore.Level)
	err := logsLevel.Set(cfg.GetMinimalLogLevel())
	if err != nil {
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
		return nil, err
	}

	cores[0] = defaultLogger.Core()

	return &Service{
		cfg:           cfg,
		defaultLogger: defaultLogger,
		cores:         cores,
	}, nil
}
