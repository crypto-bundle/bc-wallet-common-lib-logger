package logger

import (
	"log"

	"go.uber.org/zap"
)

type stdLogFabric struct {
	zapMakerFunc func(named string) *zap.Logger
}

func (s *stdLogFabric) WithFields(name string, fields map[string]interface{}) *log.Logger {
	zapAnyFields := make([]zap.Field, len(fields))
	var i int

	for fieldName, fieldValue := range fields {
		zapAnyFields[i] = zap.Any(fieldName, fieldValue)
	}

	l := s.zapMakerFunc(name).With(zapAnyFields...)

	return zap.NewStdLog(l)
}
