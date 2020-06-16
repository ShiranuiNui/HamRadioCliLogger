package main

import (
	"log"
	"os"

	"github.com/ShiranuiNui/HamRadioCliLogger/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: commands.CmdList,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
