package cmd

import (
	"strings"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"github.com/squiddy/gophercises/07-cli-task-manager/internal"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		task := strings.Join(args, " ")
		db := GetDb(cmd.Context())
		return db.Update(func(t *bolt.Tx) error {
			b := t.Bucket([]byte("todos"))
			id, _ := b.NextSequence()
			return b.Put(internal.IDtoB(id), []byte(task))
		})
	},
}
