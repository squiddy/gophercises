package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"

	"reinergerecke.de/gophercises/07-cli-task-manager/internal"
)

var rootCmd = &cobra.Command{
	Use:   "tsk",
	Short: "tsk manages your todos.",
}

type contextKey int

const databaseKey contextKey = 0

func GetDb(ctx context.Context) *bolt.DB {
	return ctx.Value(databaseKey).(*bolt.DB)
}

func Execute() {
	db, err := internal.OpenDb("/tmp/my.db")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	ctx := context.WithValue(context.Background(), databaseKey, db)
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
