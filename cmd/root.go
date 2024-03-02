package cmd

import (
	"fmt"

	"github.com/EggsyOnCode/task/db"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task",
		Short: "Task is a cli tool to manage ur tasks",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(cmd.OutOrStdout(), "Welcome to Task Cli")
		},
	}
	cmd.AddCommand(addCmd(db.GetDB()))
	cmd.AddCommand(listCmd(db.GetDB()))
	cmd.AddCommand(doneCmd(db.GetDB()))
	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.

func TestingRoot() *cobra.Command {
	return RootCmd()
}
// func init() {
// 	// Here you will define your flags and configuration settings.
// 	// Cobra supports persistent flags, which, if defined here,
// 	// will be global for your application.

// 	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.task.yaml)")

// 	// Cobra also supports local flags, which will only run
// 	// when this action is called directly.
// 	RootCmd().Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
