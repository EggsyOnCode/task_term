package cmd

import (
	"fmt"
	"strconv"

	"github.com/EggsyOnCode/task/db"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
func doneCmd(dataB db.Database) *cobra.Command {
	return &cobra.Command{
		Use:   "done",
		Short: "marks a task as done!",
		Run: func(cmd *cobra.Command, args []string) {
			var ids []int

			for _, arg := range args {
				integer, error := strconv.Atoi(arg)
				if error != nil {
					fmt.Println("parse error during str to int conversion")
				}
				ids = append(ids, integer)
			}
			tasks, err := dataB.AllTasks()
			if err != nil {
				fmt.Println("Something went wrong:", err)
				return
			}

			for _, id := range ids {
				if id <= 0 || id > len(tasks) {
					fmt.Fprint(cmd.OutOrStdout(), "Invalid task number:", id)
					continue
				}
				task := tasks[id-1]
				err := dataB.DeleteTask(task.Id)
				if err != nil {
					fmt.Fprintf(cmd.OutOrStdout(), "Failed to mark \"%d\" as completed. Error: %s\n", id, err)
				} else {
					fmt.Fprintf(cmd.OutOrStdout(), "Marked \"%d\" as completed.\n", id)
				}
			}
		},
	}
}

func init() {
	rootCmd.AddCommand(doneCmd(db.GetMockDB()))

}

func TestingDone(db db.Database) *cobra.Command {
	return doneCmd(db)
}
