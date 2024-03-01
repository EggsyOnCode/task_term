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
		fmt.Print(ids)
		for _, i := range ids {
			fmt.Print(i)
			err := db.DeleteTask(i)
			if err != nil {
				panic(err)
			}
			fmt.Printf("the task with id %d has been deleted from the task list", i)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

}
