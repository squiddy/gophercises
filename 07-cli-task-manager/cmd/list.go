package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"github.com/squiddy/gophercises/07-cli-task-manager/internal"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all incomplete tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		db := GetDb(cmd.Context())
		return db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("todos"))
			c := b.Cursor()

			fmt.Println("Your tasks:")
			i := 1
			for k, v := c.First(); k != nil; k, v = c.Next() {
				var task internal.Task
				if err := json.Unmarshal(v, &task); err != nil {
					return err
				}

				if !task.Completed.IsZero() {
					continue
				}

				fmt.Printf("%d: %s\n", i, task.Title)
				i++
			}

			return nil
		})

	},
}
