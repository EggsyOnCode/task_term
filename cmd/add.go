package cmd

import (
	"fmt"
	"strings"

	"github.com/EggsyOnCode/task/db"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
func addCmd(dataB db.Database) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "adds a task to ur task list",
		Run: func(cmd *cobra.Command, args []string) {
			task := strings.Join(args, " ")
			id, err := dataB.CreateTask(task)
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "the task \"%s\"  with id %d has been added to the task list", task, id)
		},
	}
}

func init() {
	rootCmd.AddCommand(addCmd(db.GetDB()))

}

func TestingAdd(db db.Database) *cobra.Command {
	return addCmd(db)
}
