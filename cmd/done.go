package cmd

import (
	"fmt"
	"strconv"

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
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

}
