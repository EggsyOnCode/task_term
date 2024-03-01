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
		if err != nil {
			panic(err)
		}
		for i, task := range tasks {
			fmt.Printf("number : %d \t task: %s \n", i+1, task.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
