package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"

	_ "rsc.io/sqlite"
)

const migrationFilesDir string = "migrations"

func main() {
	path := flag.String("path", "", "path to database or path where to create database")
	flag.Parse()

	if *path == "" {
		fmt.Println("path is required")
		os.Exit(64)
	}

	_, err := os.Stat(*path)
	isNewDb := false
	if os.IsNotExist(err) {
		isNewDb = true
	}

	db, err := sql.Open("sqlite3", *path)
	if err != nil {
		fmt.Printf("failed create database handle: %s\n", err.Error())
		os.Exit(65)
	}

	files, err := os.ReadDir(migrationFilesDir)
	if err != nil {
		fmt.Printf("read dir migrate failed: %s\n", err)
		os.Exit(1)
	}

	if len(files) == 0 {
		fmt.Println("migrate dir empty")
		return
	}

	if !isNewDb {
		files = files[len(files)-1:]
	}

	for _, file := range files {
		buf, err := os.ReadFile(strings.Join([]string{migrationFilesDir, file.Name()}, "/"))
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
}
