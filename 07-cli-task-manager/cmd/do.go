package cmd

import (
	"encoding/json"
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
		var task internal.Task

		db := GetDb(cmd.Context())
		id, _ := strconv.Atoi(args[0])
		if err := db.Update(func(t *bolt.Tx) error {
			b := t.Bucket([]byte("todos"))
			c := b.Cursor()

			i := 1
			for k, v := c.First(); k != nil; k, v = c.Next() {
				if i == id {
					if err := json.Unmarshal(v, &task); err != nil {
						return err
					}

					task.Complete()

					data, err := json.Marshal(task)
					if err != nil {
						return err
					}
					return b.Put(k, data)
				}
				i++
			}

			return fmt.Errorf("task with id %d not found", id)
		}); err != nil {
			return err
		}

		fmt.Printf("Marked task \"%s\" as completed.\n", task.Title)
		return nil
	},
}
