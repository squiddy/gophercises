package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark task as complete",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("accepts 1 arg, received %d", len(args))
		}

		if _, err := strconv.Atoi(args[0]); err != nil {
			return fmt.Errorf("%s is not a valid index", args[0])
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("do")
	},
}
