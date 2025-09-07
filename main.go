package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	storage "github.com/MiracleNear/dictionary/sql"
	_ "rsc.io/sqlite"
)

func main() {
	mode := flag.String("mode", "insert", "insert, update or search word in dictionary")
	flag.Parse()

	switch *mode {
	case "insert", "update":
		word, definition := readInput()
		if *mode == "insert" {
			insertWord(word, definition)
		} else {
			updateWord(word, definition)
		}
	case "search":
	default:
		fmt.Printf("%s is not valid mode, please use one of the following: insert, update or search word\n", *mode)
		os.Exit(64)
	}

}

func insertWord(word, definition string) error {
	_, err := storage.Instance.Exec(storage.InsertWord, word, definition)
	if err != nil {
		fmt.Printf("insert failed: %s\n", err.Error())
		os.Exit(1)
	}

	return nil
}

func updateWord(word, definition string) error {
	_, err := storage.Instance.Exec(storage.UpdateWord, definition, word)
	if err != nil {
		fmt.Printf("update failed: %s\n", err.Error())
		os.Exit(1)
	}

	return nil
}

func readInput() (string, string) {
	input := bufio.NewScanner(os.Stdin)
	output := make([]string, 2)
	for i, info := range []string{"Write word: ", "Write definition: "} {
		fmt.Print("\033[H\033[2J")
		fmt.Print(info)
		input.Scan()
		output[i] = input.Text()
	}

	return output[0], output[1]
}
