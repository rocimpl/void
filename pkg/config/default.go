package config

var DefaultConfig = Config{
    Push:     PushConfig{
        Host:         "",
        Port:         9000,
        ReadTimeout:  30,
        WriteTimeout: 60,
        Username:     "default",
        Password:     "",
        Database:     "default",
        Debug:        false,
    },
}
