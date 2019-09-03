package config

type Config struct {
    Targets []Target
    Push    PushConfig
}

type Target struct {
    Label string
    Host  string
}

type PushConfig struct {
    Addr     string
    Username string
    Password string
}
