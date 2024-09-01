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
	"testing"
	"time"
)

func TestMapFields(t *testing.T) {
	type testCase struct {
		logLevel          slog.Level
		logMessage        string
		attributes        []slog.Attr
		expectedAttrCount uint
	}

	var testCases = []testCase{
		{
			logLevel:   slog.LevelError,
			logMessage: "some error",
			attributes: []slog.Attr{
				slog.String("big_error_tag", "big_error_value"),
				slog.Uint64("big_uint_error_tag", 100500),
				slog.Group("big_group_error_tag",
					slog.Uint64("big_group_uint_error_tag", 100501),
					slog.Uint64("big_group_uint_error_tag", 100502),
					slog.Float64("big_group_float_error_tag", 100503.222),
				),
				slog.Uint64("big_uint_error_tag_2", 100503),
				slog.Uint64("big_uint_error_tag_3", 100504),
				slog.Uint64("big_uint_error_tag_4", 100505),
				slog.Uint64("big_uint_error_tag_5", 100506),
				slog.Uint64("big_uint_error_tag_6", 100507),
				slog.Uint64("big_uint_error_tag_7", 100508),
				slog.Group("big_group_2_error_tag",
					slog.Uint64("big_group_2_uint_error_tag", 100508),
					slog.Uint64("big_group_2_uint_error_tag", 100509),
					slog.Float64("big_group_2_float_error_tag", 100509.999),
				),
				slog.Uint64("big_uint_error_tag_8", 100510),
				slog.Uint64("big_uint_error_tag_9", 100511),
			},
			expectedAttrCount: 16,
		},
	}

	for _, tCase := range testCases {
		zapFields := MapFields(tCase.attributes)
		zapFieldsCount := uint(len(zapFields))

		if zapFieldsCount != tCase.expectedAttrCount {
			t.Errorf("zap fields count not equal with expected, current: %d, expected: %d",
				zapFieldsCount, tCase.expectedAttrCount)
		}
	}
}

func TestExtractFields(t *testing.T) {
	type testCase struct {
		logLevel          slog.Level
		logMessage        string
		attributes        []slog.Attr
		expectedAttrCount uint
	}

	var testCases = []testCase{
		{
			logLevel:   slog.LevelError,
			logMessage: "some error",
			attributes: []slog.Attr{
				slog.String("big_error_tag", "big_error_value"),
				slog.Uint64("big_uint_error_tag", 100500),
				slog.Group("big_group_error_tag",
					slog.Uint64("big_group_uint_error_tag", 100501),
					slog.Uint64("big_group_uint_error_tag", 100502),
					slog.Float64("big_group_float_error_tag", 100503.222),
				),
			},
			expectedAttrCount: 5,
		},
		{
			logLevel:   slog.LevelInfo,
			logMessage: "some info",
			attributes: []slog.Attr{
				slog.String("big_info_tag", "big_info_value"),
				slog.Uint64("big_uint_info_tag", 750),
				slog.Group("big_first_group_info_tag",
					slog.Uint64("big_first_group_uint_info_tag", 751),
					slog.Uint64("big_first_group_uint_info_tag", 752),
					slog.Float64("big_first_group_float_info_tag", 753.3331),
				),
				slog.Group("big_second_group_info_tag",
					slog.Uint64("big_second_group_uint_info_tag", 751),
					slog.Uint64("big_second_group_uint_info_tag", 752),
					slog.Float64("big_second_group_float_info_tag", 753.3331),
				),
			},
			expectedAttrCount: 8,
		},
		{
			logLevel:   slog.LevelWarn,
			logMessage: "some warn",
			attributes: []slog.Attr{
				slog.String("big_info_tag", "big_info_value"),
				slog.Uint64("big_uint_info_tag", 750),
				slog.Group("big_first_group_info_tag",
					slog.Uint64("big_first_group_uint_info_tag", 751),
					slog.Group("big_third_group_info_tag",
						slog.Uint64("big_third_group_uint_info_tag", 751),
					),
				),
			},
			expectedAttrCount: 4,
		},
	}

	for _, tCase := range testCases {
		r := slog.NewRecord(time.Now(), tCase.logLevel, tCase.logMessage, 0)
		r.AddAttrs(tCase.attributes...)

		zapFields := ExtractFields(r)
		zapFieldsCount := uint(len(zapFields))

		if zapFieldsCount != tCase.expectedAttrCount {
			t.Errorf("zap fields count not equal with expected, current: %d, expected: %d",
				zapFieldsCount, tCase.expectedAttrCount)
		}
	}
}
