package models

import (
    //"database/sql"
    _ "github.com/lib/pq"
    "log"
    "github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDB(dataSourceName string) {
    var err error
    db, err = sqlx.Open("postgres", dataSourceName)
    if err != nil {
        log.Panic(err)
    }

    if err = db.Ping(); err != nil {
        log.Panic(err)
    }
}
