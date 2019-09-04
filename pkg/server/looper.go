package server

import (
    util "github.com/rocimpl/void/pkg"
    "github.com/rocimpl/void/pkg/config"
    "github.com/rocimpl/void/pkg/follow"
    "github.com/rocimpl/void/pkg/follow/disk_file"
    "github.com/rocimpl/void/pkg/parser"
    "github.com/rocimpl/void/pkg/parser/zap"
    "github.com/rocimpl/void/pkg/pusher"
    "os"
    "os/signal"
    "syscall"
    "time"
)

type Looper struct {
    targets []follow.Follow
    pusher  *pusher.Pusher
    runner  *time.Ticker
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

func New(cfg config.Config) (looper *Looper, err error) {
    looper = new(Looper)

    looper.pusher, err = pusher.InitPusher(cfg.Push)
    if err != nil {
        return nil, err
    }

    looper.targets = make([]follow.Follow, len(cfg.Targets))
    for i, target := range cfg.Targets {
        looper.targets[i], err = newFollow(target)
        if err != nil {
            return nil, err
        }
    }

    looper.runner = time.NewTicker(cfg.Period.Duration)

    return looper, nil
}

func (looper *Looper) Run() (err error) {
    quit := make(chan os.Signal)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

    for {
        select {
        case <-looper.runner.C:
            for _, target := range looper.targets {
                read, err := target.Process(0)
                if err != nil {
                    util.Errof("Fail process", err)
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

func (looper *Looper) Shutdown() {

}
