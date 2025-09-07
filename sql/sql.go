//go:generate go run ../cmd/sql/generate/generate_sql.go
package sql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "rsc.io/sqlite"
)

type config struct {
	Database struct {
		Path   string `json:"path"`
		Driver string `json:"driver"`
	} `json:"database"`
}

var Instance *sql.DB

func init() {
	file, err := os.ReadFile("config/config.json")
	if err != nil {
		fmt.Printf("Config not found on path config/config.json: %s\n", err.Error())
		os.Exit(1)
	}

	cfg := config{}
	json.Unmarshal(file, &cfg)

	Instance, err = sql.Open(cfg.Database.Driver, cfg.Database.Path)

	if err != nil {
		fmt.Printf("failed create database handle: %s\n", err.Error())
		os.Exit(1)
	}
}
