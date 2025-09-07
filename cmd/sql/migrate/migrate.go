package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"

	_ "rsc.io/sqlite"
)

const empty = ""

func main() {
	path := flag.String("path", empty, "path for database file")
	flag.Parse()

	if *path == empty {
		fmt.Println("path is required")
		os.Exit(64)
	}

	db, err := sql.Open("sqlite3", *path)
	if err != nil {
		fmt.Printf("failed create database handle: %s\n", err.Error())
		os.Exit(65)
	}

	files, err := os.ReadDir("migrations")
	if err != nil {
		fmt.Printf("read dir migrate failed: %s\n", err)
		os.Exit(1)
	}

	if len(files) == 0 {
		fmt.Println("migrate dir empty")
		return
	}

	file := files[len(files)-1]

	buf, err := os.ReadFile(strings.Join([]string{"migrate", file.Name()}, "\\"))
	if err != nil {
		fmt.Printf("failed migration file %s: %s\n", file.Name(), err.Error())
		os.Exit(1)
	}

	_, err = db.Exec(string(buf))
	if err != nil {
		fmt.Printf("failed migration with reason: %s\n", err.Error())
		os.Exit(1)
	}
}
