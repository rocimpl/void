package config

type Config struct {
    Targets  []Target
    Push     PushConfig
    LockFile string
    Period   Duration
}

type Target struct {
    Label        []string
    Source       string
    Follow       string
    FollowParams map[string]string
    Parser       string
    ParserParams map[string]string
}

type PushConfig struct {
    Addr     string
    Port     int
    Username string
    Password string
    Debug    bool
}
