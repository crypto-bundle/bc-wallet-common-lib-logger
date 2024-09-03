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

package log

import (
	"log"

	"go.uber.org/zap"
)

type StdFields map[string]interface{}

type stdLogFabric struct {
	zapNamedMakerFunc   func(named string) *zap.Logger
	zapUnNamedMakerFunc func() *zap.Logger
}

func (s *stdLogFabric) GetLogger() *log.Logger {
	return zap.NewStdLog(s.zapUnNamedMakerFunc())
}

func (s *stdLogFabric) WithFields(fields StdFields) *log.Logger {
	zapLogger := s.zapUnNamedMakerFunc()
	zapAnyFields := make([]zap.Field, len(fields))

	var i int

	for fieldName, fieldValue := range fields {
		zapAnyFields[i] = zap.Any(fieldName, fieldValue)
	}

	return zap.NewStdLog(zapLogger.With(zapAnyFields...))
}

func (s *stdLogFabric) NamedWithFields(name string, fields StdFields) *log.Logger {
	zapAnyFields := make([]zap.Field, len(fields))
	var i int

	for fieldName, fieldValue := range fields {
		zapAnyFields[i] = zap.Any(fieldName, fieldValue)
	}

	l := s.zapNamedMakerFunc(name).With(zapAnyFields...)

	return zap.NewStdLog(l)
}

type singleNameStdLogFabric struct {
	zapMakerFunc func(named string) *zap.Logger
	name         string
}

func (s *singleNameStdLogFabric) WithFields(fields StdFields) *log.Logger {
	zapAnyFields := make([]zap.Field, len(fields))
	var i int

	for fieldName, fieldValue := range fields {
		zapAnyFields[i] = zap.Any(fieldName, fieldValue)
	}

	l := s.zapMakerFunc(s.name).With(zapAnyFields...)

	return zap.NewStdLog(l)
}

func NewStdLogMaker(zapLogger *zap.Logger) *stdLogFabric {
	return &stdLogFabric{
		zapNamedMakerFunc: zapLogger.Named,
		zapUnNamedMakerFunc: func() *zap.Logger {
			return zapLogger.With()
		},
	}
}

func NewNamedStdLogMaker(zapLogger *zap.Logger, name string) *singleNameStdLogFabric {
	return &singleNameStdLogFabric{
		zapMakerFunc: zapLogger.Named,
		name:         name,
	}
}
