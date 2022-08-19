package cmd

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"log"
	"myclass/src/config"
	"os"
)

func migrateUp() *cobra.Command {
	return &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			db, _ := config.GetConfig().Db.DB()

			driver, err := postgres.WithInstance(db, &postgres.Config{})
			if err != nil {
				log.Fatalln(err)
			}

			dir, err := os.Getwd()
			if err != nil {
				log.Fatalln(err)
			}
			path := "file://" + dir + "/src/database/postgres"

			m, err := migrate.NewWithDatabaseInstance(path, "postgres", driver)
			if err != nil {
				log.Fatalln(err)
			}

			err = m.Up()
			if err != nil {
				log.Fatalln(err)
			}

		},
	}
}

func migrateDown() *cobra.Command {
	return &cobra.Command{
		Use: "migrate-down",
		Run: func(cmd *cobra.Command, args []string) {
			db, _ := config.GetConfig().Db.DB()

			driver, err := postgres.WithInstance(db, &postgres.Config{})
			if err != nil {
				log.Fatalln(err)
			}

			dir, err := os.Getwd()
			if err != nil {
				log.Fatalln(err)
			}
			path := "file://" + dir + "/src/database/postgres"

			m, err := migrate.NewWithDatabaseInstance(path, "postgres", driver)
			if err != nil {
				log.Fatalln(err)
			}

			err = m.Down()
			if err != nil {
				log.Fatalln(err)
			}

		},
	}
}

func migrateRefresh() *cobra.Command {
	return &cobra.Command{
		Use: "migrate-refresh",
		Run: func(cmd *cobra.Command, args []string) {
			db, _ := config.GetConfig().Db.DB()

			driver, err := postgres.WithInstance(db, &postgres.Config{})
			if err != nil {
				log.Fatalln(err)
			}

			dir, err := os.Getwd()
			if err != nil {
				log.Fatalln(err)
			}
			path := "file://" + dir + "/src/database/postgres"

			m, err := migrate.NewWithDatabaseInstance(path, "postgres", driver)
			if err != nil {
				log.Fatalln(err)
			}

			err = m.Drop()
			err = m.Up()
			if err != nil {
				log.Fatalln(err)
			}

		},
	}
}
