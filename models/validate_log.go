package models

import (
    "errors"
    "strings"
)

var allowedLevels = map[string]bool{
    "debug": true,
    "info":  true,
    "warn":  true,
    "error": true,
}

// Validate checks whether the log is valid.
func (l *Log) Validate() error {
    if l.Timestamp.IsZero(){
        return errors.New("invalid timestamp")
    }

    if _, ok := allowedLevels[strings.ToLower(l.Level)]; !ok {
        return errors.New("invalid log level")
    }

    if strings.TrimSpace(l.Message) == "" {
        return errors.New("message cannot be empty")
    }

    if strings.TrimSpace(l.Source) == "" {
        return errors.New("source cannot be empty")
    }

    return nil
}
