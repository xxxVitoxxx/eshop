package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xxxVitoxxx/eshop/pkg/config"
	"github.com/xxxVitoxxx/eshop/pkg/storage/conn"
	"github.com/xxxVitoxxx/eshop/pkg/storage/mysqldb"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate",
	Long:  "migrate",
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.Load(); err != nil {
			panic(err)
		}

		db := conn.ConnectToMySQL()
		if err := db.AutoMigrate(mysqldb.Condition{}); err != nil {
			panic(err)
		}
	},
}
