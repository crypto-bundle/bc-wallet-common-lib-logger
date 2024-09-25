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

package zap

import "go.uber.org/zap"

func MakeZapFields(fields ...any) []zap.Field {
	return makeZapFields(fields...)
}

func makeZapFields(fields ...any) []zap.Field {
	fieldsCount := len(fields)
	if fieldsCount == 0 {
		return nil
	}

	isEven := fieldsCount%2 == 0
	if isEven {
		return makeFieldsForEven(fieldsCount, fields...)
	}

	return makeFieldsForNonEven(fieldsCount, fields...)
}

func makeFieldsForEven(fieldsCount int,
	fields ...any,
) []zap.Field {
	zapFields := make([]zap.Field, fieldsCount/2)

	for i, j := 0, 0; i != fieldsCount; i, j = i+2, j+1 {
		fieldName, isString := fields[i].(string)
		if !isString {
			fieldName = BadLoggerKeyName
		}

		zapFields[j] = zap.Any(fieldName, fields[i+1])
	}

	return zapFields
}

func makeFieldsForNonEven(fieldsCount int,
	fields ...any,
) []zap.Field {
	var sliceLengthCounter int
	zapFields := make([]zap.Field, (fieldsCount/2)+1)

	for index := 0; index != fieldsCount; index += 2 {
		if index+1 >= fieldsCount {
			zapFields[sliceLengthCounter] = zap.Any(BadLoggerKeyName, fields[index])
			sliceLengthCounter++

			break
		}

		fieldName, isString := fields[index].(string)
		if !isString {
			fieldName = BadLoggerKeyName
		}

		zapFields[sliceLengthCounter] = zap.Any(fieldName, fields[index+1])
		sliceLengthCounter++
	}

	return zapFields[:sliceLengthCounter]
}
