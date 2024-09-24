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

package slog

import (
	"context"
	"log/slog"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option struct {
	// log level (default: debug)
	Level slog.Leveler

	// optional: zap logger (default: zap.L())

	// optional: see slog.HandlerOptions
	AddSource   bool
	ReplaceAttr func(groups []string, a slog.Attr) slog.Attr
}

func NewZapHandler(attrs []slog.Attr, zapLogger *zap.Logger) slog.Handler {
	if zapLogger == nil {
		// should be selected lazily ?
		zapLogger = zap.L()
	}

	return &ZapHandler{
		Logger: zapLogger,
		attrs:  attrs,
		groups: []string{},
	}
}

var _ slog.Handler = (*ZapHandler)(nil)

type ZapHandler struct {
	Logger *zap.Logger
	option Option
	attrs  []slog.Attr
	groups []string
}

func (h *ZapHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.option.Level.Level()
}

func (h *ZapHandler) Handle(_ context.Context, record slog.Record) error {
	level := extractLoggerLevel(record.Level)

	fields := extractFields(record)

	checked := h.Logger.Check(level, record.Message)

	switch true {
	case checked == nil:
		fallthrough
	default:
		h.Logger.Log(level, record.Message, fields...)

	case checked != nil && h.option.AddSource:
		frame, _ := runtime.CallersFrames([]uintptr{record.PC}).Next()
		checked.Caller = zapcore.NewEntryCaller(0, frame.File, frame.Line, true)
		checked.Stack = "" //@TODO

		checked.Write(fields...)

	case checked != nil && !h.option.AddSource:
		checked.Caller = zapcore.EntryCaller{}
		checked.Stack = ""

		checked.Write(fields...)
	}

	return nil
}

func (h *ZapHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ZapHandler{
		option: h.option,
		attrs:  append(h.attrs, attrs...),
		groups: h.groups,
	}
}

func (h *ZapHandler) WithGroup(name string) slog.Handler {
	// https://cs.opensource.google/go/x/exp/+/46b07846:slog/handler.go;l=247
	if name == "" {
		return h
	}

	return &ZapHandler{
		option: h.option,
		attrs:  h.attrs,
		groups: append(h.groups, name),
	}
}
