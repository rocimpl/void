package pusher

import (
    "database/sql"
    "github.com/rocimpl/void/pkg/config"
    "github.com/rocimpl/void/pkg/types"
)

type Pusher struct {
    db *sql.DB
}

func InitPusher(p *config.PushConfig) *Pusher {
    return &Pusher{}
}

func (p *Pusher) PushSnapshot(snapshot []types.LogFormat) error {
    if !(len(snapshot) > 0) {
        return nil
    }

    tx, err := p.db.Begin()
    if err != nil {
        return err
    }

    err = tx.Commit()
    if err != nil {
        return err
    }

    return nil
}
