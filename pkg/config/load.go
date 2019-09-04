package config

import (
    "github.com/BurntSushi/toml"
)

func LoadConfig(filename string, cfg *Config) error {
    *cfg = DefaultConfig
    _, err := toml.DecodeFile(filename, &cfg)
    if err != nil {
        return err
    }

    return nil
}
