package main

import (
    "github.com/rocimpl/void/pkg"
    "github.com/rocimpl/void/pkg/config"
    "github.com/rocimpl/void/pkg/server"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "os"
)

func init() {
    logger := zap.New(
        zapcore.NewCore(
            zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
            os.Stdout,
            zap.NewAtomicLevel(),
        ),
    )

    util.InitLogger(logger)
}

func main() {
    var cfg config.Config

    if err := config.LoadConfig("example.config.toml", &cfg); err != nil {
        util.Errof("Fail loading config", err)
        os.Exit(1)
    }

    looper, err := server.New(cfg)
    if err != nil {
        util.Errof("Fail init server", err)
        os.Exit(1)
    }

    if err := looper.Run(); err != nil {
        util.Errof("Fail run server", err)
        os.Exit(1)
    }

    looper.Shutdown()
}
