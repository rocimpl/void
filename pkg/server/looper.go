package server

import (
    "github.com/rocimpl/void/pkg/config"
    "github.com/rocimpl/void/pkg/follow"
    "github.com/rocimpl/void/pkg/follow/disk_file"
    "github.com/rocimpl/void/pkg/parser"
    "github.com/rocimpl/void/pkg/parser/zap"
    "os"
    "os/signal"
    "syscall"
    "time"
)

type Looper struct {
    config  config.Config
    targets []follow.Follow
}

func newFollow(target config.Target) (follow.Follow, error) {
    belt, err := newParser(target)
    if err != nil {
        return nil, err
    }

    switch target.Follow {
    case "disk_file":
        return disk_file_follow.NewDiskFileFollow(target.FollowParams, belt)
    default:
        return nil, follow.ErrUnknownFollow
    }
}

func newParser(target config.Target) (parser.Parser, error) {
    switch target.Parser {
    case "zap":
        return zap.NewZapParse(target.ParserParams)
    default:
        return nil, parser.ErrUnknownParser
    }
}

func New(cfg config.Config) (lp *Looper, err error) {
    lp = &Looper{
        config: cfg,
    }

    lp.targets = make([]follow.Follow, len(lp.config.Targets))
    for i, target := range lp.config.Targets {
        lp.targets[i], err = newFollow(target)
        if err != nil {
            return nil, err
        }
    }

    return lp, nil
}

func (lp *Looper) Run() (err error) {
    quit := make(chan os.Signal)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    runner := time.NewTicker(lp.config.Period.Duration)

    for {
        select {
        case <-runner.C:
            for _, target := range lp.targets {
                read, err := target.Process(0)
                if err != nil {
                    // Log error
                    continue
                }

                if !(len(read) > 0) {
                    continue
                }

                //read
            }
        case <-quit:
            return nil
        }
    }
}

func (lp *Looper) Shutdown() {

}
