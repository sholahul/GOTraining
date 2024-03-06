package cmd

import (
	"OAUTH2/migrations"
	"OAUTH2/pkg/db"
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/uptrace/bun/migrate"
)

var (
	MigrateName    []string
	MigrateCommand string
)

func init() {
	autoMigrate.Flags().StringSliceVar(&MigrateName, "name", nil, "Name of the migration")
	autoMigrate.Flags().StringVar(&MigrateCommand, "command", "", "Migration command to Execute")
	rootCmd.AddCommand(autoMigrate)
}

var autoMigrate = &cobra.Command{
	Use:   "db",
	Short: "Start db",
	Run: func(cmd *cobra.Command, args []string) {
		switch MigrateCommand {
		case "init":
			initMigration()
		case "create_go":
			createGoMigration()
		case "migrate":
			doMigration()
		case "rollback":
			rollbackMigration()
		default:
			fmt.Println("command not found")
		}
	},
}

func initMigration() {
	ctx := context.Background()
	newDB := db.GetDB()
	mg := migrate.NewMigrator(newDB, migrations.Migrations)
	err := mg.Init(ctx)

	if err != nil {
		panic(err)
	}
	fmt.Println("Inisialisasi Berhasil")
}

func rollbackMigration() {
	fmt.Println("rollback the last migration group")
	ctx := context.Background()
	newDB := db.GetDB()
	mg := migrate.NewMigrator(newDB, migrations.Migrations)
	mg.Unlock(ctx)
	if err := mg.Lock(ctx); err != nil {
		panic(err)
	}
	defer mg.Unlock(ctx)
	group, err := mg.Rollback(ctx)

	if err != nil {
		panic(err)
	}

	if group.IsZero() {
		fmt.Printf("there are no groups to rollback \n")
	}
	fmt.Printf("rolled back %s \n", group)
}

func createGoMigration() {
	fmt.Println("create go migrations")
	ctx := context.Background()
	newDB := db.GetDB()
	mg := migrate.NewMigrator(newDB, migrations.Migrations)
	name := strings.Join(MigrateName, "_")
	mf, err := mg.CreateGoMigration(ctx, name)
	if err != nil {
		panic(err)
	}
	fmt.Printf("create migration %s (%s) \n", mf.Name, mf.Path)
}

func doMigration() {
	fmt.Println("Start Migrate...")
	ctx := context.Background()
	newDB := db.GetDB()
	mg := migrate.NewMigrator(newDB, migrations.Migrations)
	if err := mg.Lock(ctx); err != nil {
		panic(err)
	}
	defer mg.Unlock(ctx)
	group, err := mg.Migrate(ctx)
	if err != nil {
		panic(err)
	}
	if group.IsZero() {
		fmt.Printf("there are no new migrations to run (database is up to date)\n")
	}
}
