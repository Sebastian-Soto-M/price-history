package main

import (
	"os"

	"github.com/Sebastian-Soto-M/price-history/config"
)

func main() {
	app := config.NewApp()
	app.Run(os.Args)
}
