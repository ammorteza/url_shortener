package main

import (
	"github.com/ammorteza/urlShortener/src/db"
	"github.com/ammorteza/urlShortener/src/db/migrations"
	"github.com/ammorteza/urlShortener/src/interfaces"
	"github.com/urfave/cli"
	"log"
	"os"
)

var app = cli.NewApp()
var dbConnection interfaces.DbConnection

func initDnConnection()  {
	dbConnection = db.MysqlConnection{}
}

func info(){
	app.Name = "url shortener cli"
	app.Usage = "use it for run a few commands"
	app.Author = "Morteza Amzajerdi"
	app.Version = "1.0.0"
}

func commands()  {
	app.Commands = []cli.Command{
		{
			Name: 	"db:migrate",
			Usage: 	"migrate all database migration files",
			Action: func(c *cli.Context) {
				var migration interfaces.MigrationCliInterface = migrations.MigrationCli{}
				migration.Make(dbConnection)
			},
		},
		{
			Name: 	"db:reset",
			Usage: 	"reset all database tables",
			Action: func(c *cli.Context) {
				var migration interfaces.MigrationCliInterface = migrations.MigrationCli{}
				migration.Reset(dbConnection)
			},
		},
	}
}

func main()  {
	info()
	commands()
	initDnConnection()

	err := app.Run(os.Args)
	if err != nil{
		log.Fatal(err)
	}
}
