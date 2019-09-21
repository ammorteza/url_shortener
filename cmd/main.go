package main

import (
	"github.com/ammorteza/url_shortener/db"
	"github.com/ammorteza/url_shortener/db/migrations"
	"github.com/ammorteza/url_shortener/router"
	"github.com/urfave/cli"
	"log"
	"os"
)

var app = cli.NewApp()
var dbConnection db.DbConnection

func initDnConnection() {
	dbConnection = db.MysqlConnection{}
}

func info() {
	app.Name = "url shortener cli"
	app.Usage = "use it for run a few commands"
	app.Author = "Morteza Amzajerdi"
	app.Version = "1.1.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:  "router:start",
			Usage: "start router",
			Action: func(c *cli.Context) {
				_router := router.Router{}
				_router.Start(db.MysqlConnection{})
			},
		},
		{
			Name:  "db:migrate",
			Usage: "migrate all database migration files",
			Action: func(c *cli.Context) {
				var migration migrations.MigrationCliInterface = migrations.MigrationCli{}
				migration.Make(dbConnection)
			},
		},
		{
			Name:  "db:reset",
			Usage: "reset all database tables",
			Action: func(c *cli.Context) {
				var migration migrations.MigrationCliInterface = migrations.MigrationCli{}
				migration.Reset(dbConnection)
			},
		},
	}
}

func main() {
	info()
	commands()
	initDnConnection()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
