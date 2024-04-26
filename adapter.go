package watermilllogr

import (
	"errors"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/go-logr/logr"
)

var NotImplemented = errors.New("not implemented")

// Implements the Watermill LoggerAdapter with the logr Logger interface
type Logger struct {
	backend *logr.Logger
	fields  watermill.LogFields
}

// NewLogger returns new watermill.LoggerAdapter using passed *zap.Logger as backend.
func NewLogger(l *logr.Logger) watermill.LoggerAdapter {
	return &Logger{backend: l}
}

func (l *Logger) Error(msg string, err error, fields watermill.LogFields) {
	l.backend.Error(err, msg, l.fieldsAsKeysAndValues(fields)...)
}

func (l *Logger) Info(msg string, fields watermill.LogFields) {
	l.backend.Info(msg, l.fieldsAsKeysAndValues(fields)...)
}

func (l *Logger) Debug(msg string, fields watermill.LogFields) {
	l.backend.V(1).Info(msg, l.fieldsAsKeysAndValues(fields)...)
}

func (l *Logger) Trace(msg string, fields watermill.LogFields) {
	l.backend.V(2).Info(msg, l.fieldsAsKeysAndValues(fields)...)
}

func (l *Logger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	return &Logger{
		backend: l.backend,
		fields:  l.fields.Add(fields),
	}
}

func (l *Logger) fieldsAsKeysAndValues(fields watermill.LogFields) []any {
	var keysAndValues []any
	for k, v := range l.fields {
		keysAndValues = append(keysAndValues, k, v)
	}
	for k, v := range fields {
		keysAndValues = append(keysAndValues, k, v)
	}
	return keysAndValues
}
