package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/EggsyOnCode/task/cmd"
	"github.com/EggsyOnCode/task/db"
	homeDir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homeDir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.InitDB(dbPath)
	if db.GetDB() == nil {
		fmt.Print("the db has not been init")
	}
	if err != nil {
		panic(err)
	}
	must(cmd.RootCmd().Execute())
}
func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
