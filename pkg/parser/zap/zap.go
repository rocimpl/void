package zap

import (
    "encoding/json"
    "errors"
    "github.com/rocimpl/void/pkg/parser"
    "github.com/rocimpl/void/pkg/types"
    "time"
)

const (
    DefaultZapTimeKey       = "ts"
    DefaultZapLevelKey      = "level"
    DefaultZapNameKey       = "logger"
    DefaultZapCallerKey     = "caller"
    DefaultZapMessageKey    = "msg"
    DefaultZapStacktraceKey = "stacktrace"
)

type zapParser struct {
    timekey       string
    levelkey      string
    namekey       string
    callerkey     string
    messagekey    string
    stacktracekey string
}

func (p *zapParser) Parse(sequence []byte) (parsed types.LogFormat, err error) {
    var log map[string]interface{}
    if err := json.Unmarshal(sequence, &log); err != nil {
        return parsed, err
    }

    if log[p.timekey] == nil {
        return parsed, errors.New("bad key")
    }

    if log[p.messagekey] == nil {
        return parsed, errors.New("bad key")
    }

    if log[p.levelkey] == nil {
        return parsed, errors.New("bad key")
    }

    parsed.Time = time.Unix(0, int64(log[p.timekey].(float64)*float64(time.Second)))
    parsed.Message = log[p.messagekey].(string)

    switch log[p.levelkey].(string) {
    case "debug":
        parsed.Level = types.DebugLevel
    case "info":
        parsed.Level = types.InfoLevel
    case "warn":
        parsed.Level = types.WarningLevel
    case "error":
        parsed.Level = types.ErrorLevel
    case "dpanic":
    case "panic":
    case "fatal":
        parsed.Level = types.FatalLevel
    }

    return parsed, nil
}

func NewZapParse(params map[string]string) (parser.Parser, error) {
    return &zapParser{
        timekey:       DefaultZapTimeKey,
        levelkey:      DefaultZapLevelKey,
        namekey:       DefaultZapNameKey,
        callerkey:     DefaultZapCallerKey,
        messagekey:    DefaultZapMessageKey,
        stacktracekey: DefaultZapStacktraceKey,
    }, nil
}
