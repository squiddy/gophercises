package cmd

import (
	"encoding/json"
	"fmt"
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
		title := strings.Join(args, " ")
		task := internal.NewTask(title)
		data, err := json.Marshal(task)
		if err != nil {
			return err
		}

		db := GetDb(cmd.Context())
		if err := db.Update(func(t *bolt.Tx) error {
			b := t.Bucket([]byte("todos"))
			id, _ := b.NextSequence()
			return b.Put(internal.IDtoB(id), data)
		}); err != nil {
			return err
		}

		fmt.Printf("Added \"%s\" to the tasks.\n", title)

		return nil
	},
}
