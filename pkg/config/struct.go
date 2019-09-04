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
    Host         string
    TableName    string
    Port         int
    ReadTimeout  int
    WriteTimeout int
    Username     string
    Password     string
    Database     string
    Debug        bool
}
