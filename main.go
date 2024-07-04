package main

import (
	"fmt"
	"go-minimalists/app"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	cli := &cli.App{
		Name:    "Go Minimalists",
		Usage:   "make an explosive entrance",
		Suggest: true,
		Commands: []*cli.Command{
			{
				Name:  "serve",
				Usage: "Serve the application",
				Action: func(ctx *cli.Context) error {
					app := app.NewApp(ctx.Context).WithDBConnect()
					app.Serve()

					return nil
				},
			},
			{
				Name:  "route:list",
				Usage: "List all registered routes",
				Action: func(ctx *cli.Context) error {
					app := app.NewApp(ctx.Context)

					return app.RouteList()
				},
			},
			{
				Name:  "migrate",
				Usage: "Run the database migrations",
				Action: func(ctx *cli.Context) (err error) {
					app := app.NewApp(ctx.Context).WithDBConnect()

					if err := app.DB().MySql.AutoMigrate(migrations...); err != nil {
						fmt.Printf("Migrate err: %s\n", err.Error())
						return err
					}

					fmt.Println("Migrate OK")

					return nil
				},
			},
		},
	}

	if err := cli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
