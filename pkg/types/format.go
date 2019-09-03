package types

import "time"

type LogFormat struct {
    Time    time.Time
    Level   LogLevel
    Message string
    Trace   []string
}
