package cmd

import (
	"fmt"

	"github.com/EggsyOnCode/task/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all the tasks currently in the task list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		tasks, err := db.AllTasks()
		if err!= nil{
			panic(err)
		}
		for id, value := range tasks{
			fmt.Printf("id : %d \t task: %s \n", id, value.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
