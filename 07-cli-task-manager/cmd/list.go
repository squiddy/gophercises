package cmd

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
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
				fmt.Printf("%d: %s\n", i, v)
				i++
			}

			return nil
		})

	},
}
