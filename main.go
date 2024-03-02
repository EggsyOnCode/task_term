package main

import (
	"path/filepath"
	"github.com/EggsyOnCode/task/cmd"
	"github.com/EggsyOnCode/task/db"
	homeDir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homeDir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.InitDB(dbPath)
	if err != nil {
		panic(err)
	}
	cmd.Execute()
}
