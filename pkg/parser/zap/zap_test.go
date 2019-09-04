package zap_test

import (
    "bufio"
    "bytes"
    "fmt"
    "testing"
    "time"
)

var testZap = []byte(`{"level":"info","ts":1566390893.1826751,"msg":"78 Madison Ave,\nHoonah, ID, 83725","int":4751997750760398084}
{"level":"info","ts":1566390893.1826751,"msg":"72 Washington Circle,\nLucien, MI, 59584","int":7504504064263669287}
{"level":"info","ts":1566390893.1826751,"msg":"73 Washington Ter,\nCampden, KS, 76103","int":1976235410884491574}
{"level":"info","ts":1566390893.1826751,"msg":"54 Adams Pl,\nDerby Center, TN, 57645","int":3510942875414458836}`)

func TestZapParse(t *testing.T) {
    zap := NewZapParse(nil)

    r := bytes.NewBuffer(testZap)
    s := bufio.NewScanner(r)
    for s.Scan() {
        if err := zap.Parse(s.Bytes()); err != nil {
            t.Fatal(err)
        }
    }

    for _, shot := range zap.Snapshot() {
        fmt.Println(shot)
    }
}

func TestRace(t *testing.T) {
    var snapshot []string
    var buffer []string

    for i := 0; i < 30; i++ {
        go func() {
           for {
               buffer = append(buffer, "123")
           }
        }()
    }

    for i := 0; i < 30; i++ {
        go func() {
            for {
                buffer, snapshot = make([]string, 0, 32), buffer
                time.Sleep(time.Millisecond*2)
            }
        }()
    }

    time.Sleep(time.Second)
    fmt.Println(snapshot)
}
