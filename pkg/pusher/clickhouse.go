package pusher

import (
    "github.com/rocimpl/void/pkg/config"
    "github.com/rocimpl/void/pkg/types"
)

type Pusher struct {
}

func InitPusher(p *config.PushConfig) *Pusher {
    return &Pusher{}
}

func (p *Pusher) PushSnapshot(snapshot []types.LogFormat) error {
    if !(len(snapshot) > 0) {
        return nil
    }


}
