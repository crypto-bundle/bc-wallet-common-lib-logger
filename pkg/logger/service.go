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
	"log"
	"log/slog"

	cllog "github.com/crypto-bundle/bc-wallet-common-lib-logger/pkg/logger/log"
	clslog "github.com/crypto-bundle/bc-wallet-common-lib-logger/pkg/logger/slog"
	clzap "github.com/crypto-bundle/bc-wallet-common-lib-logger/pkg/logger/zap"

	"go.uber.org/zap"
)

type Service struct {
	cfg configManager

	zapLogEntryBuilderSvc  zapLogEntryService
	stdLogEntryBuilderSvc  stdLogEntryService
	slogLogEntryBuilderSvc slogLogEntryService
}

func (s *Service) NewStdLogger(named string, fields ...any) *log.Logger {
	return s.stdLogEntryBuilderSvc.NewLoggerEntry(named, fields...)
}

func (s *Service) NewSlogEntry(named string, fields ...any) *slog.Logger {
	return s.slogLogEntryBuilderSvc.NewLoggerEntry(named, fields...)
}

func (s *Service) NewLoggerEntry(named string, fields ...any) *zap.Logger {
	return s.zapLogEntryBuilderSvc.NewLoggerEntry(named, fields...)
}

func NewService(cfg configManager) (*Service, error) {
	var fields = []any{
		HostnameFieldTag, cfg.GetHostName(),
		StageNameTag, cfg.GetStageName(),
		ApplicationPID, cfg.GetApplicationPID(),
		EnvironmentNameTag, cfg.GetEnvironmentName(),
	}

	isDevOrLocal := cfg.IsDev() || cfg.IsLocal()
	buildInfoEnabled := isDevOrLocal && cfg.GetSkipBuildInfo()
	if buildInfoEnabled {
		fields = append(fields, []any{
			SVCReleaseTag, cfg.GetReleaseTag(),
			SVCCommitShortID, cfg.GetShortCommitID(),
			SVCCommitID, cfg.GetCommitID(),
			BuildNumberTag, cfg.GetBuildNumber(),
			BuildDateTag, cfg.GetBuildDate(),
			BuildDateTimestampTag, uint64(cfg.GetBuildDateTS()),
		})
	}

	zapLoggerEntryBuilder, err := clzap.NewService(cfg, fields...)
	if err != nil {
		return nil, err
	}

	slogBuilderSvc := clslog.NewSLogMaker(zapLoggerEntryBuilder)
	stdLogBuilderSvc := cllog.NewStdLogMaker(zapLoggerEntryBuilder)

	return &Service{
		cfg: cfg,

		zapLogEntryBuilderSvc:  zapLoggerEntryBuilder,
		stdLogEntryBuilderSvc:  stdLogBuilderSvc,
		slogLogEntryBuilderSvc: slogBuilderSvc,
	}, nil
}
