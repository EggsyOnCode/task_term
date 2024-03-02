package cmd

import (
	"fmt"
	"github.com/EggsyOnCode/task/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
func listCmd(dataB db.Database) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "lists all the tasks currently in the task list",
		Run: func(cmd *cobra.Command, args []string) {
			tasks, err := dataB.AllTasks()
			if err != nil {
				panic(err)
			}
			if len(tasks) == 0 {
				fmt.Println("Hurrah! all the tasks completed!\n Now u may sleep soundly!")
			}
			for i, task := range tasks {
				fmt.Fprintf(cmd.OutOrStdout(), "number : %d \t task: %s \n", i+1, task.Value)
			}
		},
	}
}

// func init() { rootCmd.AddCommand(listCmd(db.GetDB()))

// }

func TestingList(db db.Database) *cobra.Command {
	return listCmd(db)
}
