package parser

import (
    "encoding/json"
    "errors"
    "github.com/rocimpl/void/pkg/config"
    "github.com/rocimpl/void/pkg/types"
    "sync"
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
    accumulator   []types.LogFormat
    mx sync.Mutex
}

func (p *zapParser) Snapshot() (snapshot []types.LogFormat) {
    p.mx.Lock()
    defer p.mx.Unlock()
    p.accumulator, snapshot = make([]types.LogFormat, 0, 32), p.accumulator
    return snapshot
}

func (p *zapParser) Parse(sequence []byte) error {
    var log map[string]interface{}
    if err := json.Unmarshal(sequence, &log); err != nil {
        return err
    }

    if log[p.timekey] == nil {
        return errors.New("bad key")
    }

    if log[p.messagekey] == nil {
        return errors.New("bad key")
    }

    if log[p.levelkey] == nil {
        return errors.New("bad key")
    }

    var pure = types.LogFormat{
       Time:    time.Unix(0, int64(log[p.timekey].(float64)*float64(time.Second))),
       Message: log[p.messagekey].(string),
    }

    switch log[p.levelkey].(string) {
    case "debug":
        pure.Level = types.DebugLevel
    case "info":
        pure.Level = types.InfoLevel
    case "warn":
        pure.Level = types.WarningLevel
    case "error":
        pure.Level = types.ErrorLevel
    case "dpanic":
    case "panic":
    case "fatal":
        pure.Level = types.FatalLevel
    }

    p.mx.Lock()
    defer p.mx.Unlock()
    p.accumulator = append(p.accumulator, pure)
    return nil
}

func NewZapParse(target *config.Target) Parser {
    return &zapParser{
        timekey:       DefaultZapTimeKey,
        levelkey:      DefaultZapLevelKey,
        namekey:       DefaultZapNameKey,
        callerkey:     DefaultZapCallerKey,
        messagekey:    DefaultZapMessageKey,
        stacktracekey: DefaultZapStacktraceKey,
        accumulator:   make([]types.LogFormat, 0, 32),
    }
}
