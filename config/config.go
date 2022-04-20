package config

import (
	"fmt"

	"github.com/Sebastian-Soto-M/price-history/backend"
	"github.com/Sebastian-Soto-M/price-history/backend/store"
	"github.com/Sebastian-Soto-M/price-history/models"
	"github.com/urfave/cli/v2"
)

func addCommand() *cli.Command {
	return &cli.Command{
		Name:  "add",
		Usage: "track a new item",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "store",
				Value:   "amazon",
				Aliases: []string{"s"},
			},
		},
		Action: func(ctx *cli.Context) error {
			store := *models.StoreEnum(ctx.String("store")).Store()
			fmt.Println(store)
			return nil
		},
	}
}

func listCommand() *cli.Command {
	stores := &cli.Command{
		Name:  "stores",
		Usage: "lists stores you are able to use",
		Action: func(ctx *cli.Context) error {
			stores := storebe.GetAll()
			fmt.Println("Total:", len(stores))
			for _, store := range stores {
				fmt.Println(store.Name)
			}
			return nil
		},
	}

	items := &cli.Command{
		Name:  "items",
		Usage: "lists items you are currently tracking",
		Action: func(ctx *cli.Context) error {
			return nil
		},
	}

	return &cli.Command{
		Name:        "list",
		Usage:       "List the information you have available",
		HideHelp:    true,
		Subcommands: []*cli.Command{stores, items},
	}
}

func fetchCommand() *cli.Command {
	return &cli.Command{
		Name:  "fetch",
		Usage: "gets updated information for your tracked items",
		Action: func(c *cli.Context) error {
			fmt.Printf("the next example")
			return nil
		},
	}
}

func cleanCommand() *cli.Command {
	return &cli.Command{
		Name:  "clean",
		Usage: "removes all data from the database",
		Action: func(ctx *cli.Context) error {
			backend.CleanDatabase()
			return nil
		},
	}
}

func reportCommand() *cli.Command {
	return &cli.Command{
		Name:  "report",
		Usage: "Shows a report with data from all your tracked items",
		Action: func(c *cli.Context) error {
			fmt.Printf("i like to describe things")
			return nil
		},
	}
}

func NewApp() *cli.App {
	return &cli.App{
		Name:                 "sale-history",
		Usage:                "tracks the price history of stuff online",
		Description:          "A tool that tracks sales in your favorite online stores",
		Version:              "0.0.1",
		HideVersion:          true,
		EnableBashCompletion: true,
		HideHelpCommand:      true,
		Commands: []*cli.Command{
			addCommand(),
			cleanCommand(),
			fetchCommand(),
			listCommand(),
			reportCommand(),
		},

		Authors: []*cli.Author{{Name: "Sebastian Soto M", Email: "s.m.sebastian.n@gmail.com"}},
	}
}
