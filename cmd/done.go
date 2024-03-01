package cmd

import (
	"fmt"
	"strconv"

	"github.com/EggsyOnCode/task/db"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
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
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Id)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", id)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

}
