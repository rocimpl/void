package types

type LogLevel int8

const (
    DebugLevel LogLevel = iota
    InfoLevel
    WarningLevel
    ErrorLevel
    FatalLevel
)
