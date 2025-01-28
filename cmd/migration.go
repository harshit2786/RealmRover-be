package cmd

import (
	"fmt"
	"realmrovers/config"
	"realmrovers/db"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use: "migrate",
	Short: "Run database migrations",
	Long:  "This command applies all database migrations using GORM.",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfig()
		dbc := db.ConnectDb(cfg)
		db.MigrateDB(dbc)
		fmt.Println("Database migrated successfully!")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}