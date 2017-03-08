package config

import (
	"database/sql"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
	"github.com/mantishK/galore/log"
)

var DB *sql.DB

func init() {
	dbUserName, ok := GetString("pg_uname")
	if !ok {
		log.Err("Config values missing")
		panic("Config values missing, unable to start the server")
	}
	dbPass, ok := GetString("pg_pass")
	if !ok {
		log.Err("Config values missing")
		panic("Config values missing, unable to start the server")
	}

	dbIp, ok := GetString("pg_ip")
	if !ok {
		log.Err("Config values missing")
		panic("Config values missing, unable to start the server")
	}

	dbPortNo, ok := GetInt("pg_port")
	if !ok {
		log.Err("Config values missing")
		panic("Config values missing, unable to start the server")
	}

	dbName, ok := GetString("pg_name")
	if !ok {
		log.Err("Config values missing")
		panic("Config values missing, unable to start the server")
	}
	if dbPass != "" {
		dbPass = ":" + dbPass
	} else {
		dbPass = ":"
	}
	var err error
	dbStringSlice := []string{"postgres://", dbUserName, dbPass, "@", dbIp, ":", strconv.Itoa(dbPortNo), "/", dbName, "?sslmode=disable"}
	DB, err = sql.Open("postgres", strings.Join(dbStringSlice, ""))
	if err != nil {
		log.Err("Unable to connect to DB")
		panic("Unable to connect to DB")
	}

}
