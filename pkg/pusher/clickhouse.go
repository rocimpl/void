package pusher

import (
    "database/sql"
    "fmt"
    _ "github.com/kshvakov/clickhouse"
    util "github.com/rocimpl/void/pkg"
    "github.com/rocimpl/void/pkg/config"
    "github.com/rocimpl/void/pkg/types"
    "strconv"
)

const baseSQLInsert = `INSERT INTO %s (host, label, level, message) VALUES (?, ?, ?, ?)`

type Pusher struct {
    db        *sql.DB
    sqlInsert string
}

func InitPusher(cfg config.PushConfig) (p *Pusher, err error) {
    p = &Pusher{
        sqlInsert: fmt.Sprintf(baseSQLInsert, cfg.TableName),
    }

    dataSourceName := fmt.Sprintf(
        "tcp://%s:%d?username=%s&password=%s&database=%s&read_timeout=%d&write_timeout=%d&debug=%t",
        cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database, cfg.ReadTimeout, cfg.WriteTimeout, cfg.Debug,
    )

    p.db, err = sql.Open("clickhouse", dataSourceName)
    if err != nil {
        return nil, err
    }

    return p, nil
}

func (p *Pusher) PushSnapshot(snapshot []types.LogFormat) error {
    if !(len(snapshot) > 0) {
        return nil
    }

    tx, err := p.db.Begin()
    if err != nil {
        return err
    }

    stmt, err := tx.Prepare(p.sqlInsert)
    if err != nil {
        return err
    }

    err = p.process(stmt, snapshot)
    if err := stmt.Close(); err != nil {
        return err
    }

    if err != nil {
        if err := tx.Rollback(); err != nil {
            util.Errof("Fail rollback transaction", err)
        }

        return err
    }

    err = tx.Commit()
    if err != nil {
        return err
    }

    return nil
}

func (p *Pusher) process(stmt *sql.Stmt, snapshot []types.LogFormat) error {
    for _, s := range snapshot {
        if _, err := stmt.Exec("1", "testx", "debug", strconv.Itoa(i)); err != nil {
            return err
        }
    }

    return nil
}
