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
	"log/slog"

	"go.uber.org/zap"
)

func getFieldMapCallback(kind slog.Kind) func(attr slog.Attr) zap.Field {
	switch kind {
	case slog.KindBool:
		return func(attr slog.Attr) zap.Field {
			return zap.Bool(attr.Key, attr.Value.Bool())
		}

	case slog.KindDuration:
		return func(attr slog.Attr) zap.Field {
			return zap.Duration(attr.Key, attr.Value.Duration())
		}

	case slog.KindString:
		return func(attr slog.Attr) zap.Field {
			return zap.String(attr.Key, attr.Value.String())
		}

	case slog.KindTime:
		return func(attr slog.Attr) zap.Field {
			return zap.Time(attr.Key, attr.Value.Time())
		}

	case slog.KindInt64:
		return func(attr slog.Attr) zap.Field {
			return zap.Int64(attr.Key, attr.Value.Int64())
		}

	case slog.KindUint64:
		return func(attr slog.Attr) zap.Field {
			return zap.Uint64(attr.Key, attr.Value.Uint64())
		}

	case slog.KindFloat64:
		return func(attr slog.Attr) zap.Field {
			return zap.Float64(attr.Key, attr.Value.Float64())
		}

	default:
		return nil
	}
}

func ExtractFields(record slog.Record) []zap.Field {
	zapFields := make([]zap.Field, record.NumAttrs())

	index := 0

	record.Attrs(func(attr slog.Attr) bool {
		kind := attr.Value.Kind()

		clb := getFieldMapCallback(attr.Value.Kind())
		if clb != nil {
			zapFields[index] = clb(attr)
			index++

			return true
		}

		if kind == slog.KindGroup {
			groupFields := MapFields(attr.Value.Group())

			zapFields = append(zapFields, make([]zap.Field, len(groupFields)-1)...)
			copy(zapFields[index:], groupFields)

			index += len(groupFields)
		}

		return true
	})

	return zapFields
}

func MapFields(attrs []slog.Attr) []zap.Field {
	zapFields := make([]zap.Field, len(attrs))
	index := 0

	for i, _ := range attrs {
		attr := attrs[i]
		kind := attr.Value.Kind()

		clb := getFieldMapCallback(attr.Value.Kind())
		if clb != nil {
			zapFields[index] = clb(attr)
			index++

			continue
		}

		if kind == slog.KindGroup {
			groupFields := MapFields(attr.Value.Group())

			zapFields = append(zapFields, make([]zap.Field, len(groupFields)-1)...)
			copy(zapFields[index:], groupFields)

			index += len(groupFields)
		}
	}

	return zapFields
}
