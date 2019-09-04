package main

import "github.com/BurntSushi/toml"

func main() {
    //db, err := sql.Open("clickhouse", "tcp://vps6192.mtu.immo:9000?debug=true&skip_verify=true")
    //if err != nil {
    //    panic(err)
    //}
    //
    ////err = db.Ping()
    ////if err != nil {
    ////    panic(err)
    ////}
    //
    //tx, err := db.Begin()
    //if err != nil {
    //    panic(err)
    //}
    //
    //stmt, err := tx.Prepare(`INSERT INTO high_load_test (host, label, level, message) VALUES (?, ?, ?, ?)`)
    //if err != nil {
    //    panic(err)
    //}
    //
    //for i := range make([]string, 600) {
    //    _, err := stmt.Exec("1", "testx", "debug", strconv.Itoa(i))
    //    if err != nil {
    //        panic(err)
    //    }
    //}
    //
    //err = stmt.Close()
    //if err != nil {
    //    panic(err)
    //}
    //
    //err = tx.Commit()
    //if err != nil {
    //    panic(err)
    //}
    //
    //err = db.Close()
    //if err != nil {
    //    panic(err)
    //}


}
