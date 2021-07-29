package cmd

import (
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"github.com/squiddy/gophercises/07-cli-task-manager/internal"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		db := GetDb(cmd.Context())
		id, _ := strconv.Atoi(args[0])
		return db.Update(func(t *bolt.Tx) error {
			b := t.Bucket([]byte("todos"))
			return b.Delete(internal.IDtoB(uint64(id)))
		})
	},
}
