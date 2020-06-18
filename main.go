package main

import (
	"log"
	"os"
	"os/user"

	"github.com/ShiranuiNui/HamRadioCliLogger/commands"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func main() {
	usr, _ := user.Current()
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:  "config",
			Usage: "Place of config file",
			Value: usr.HomeDir + "/.config/hamradio_logger/config.yaml",
		},
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "mycallsign",
			Usage: "Your CallSign",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "qth",
			Usage: "Your Position Infomation",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "rig",
			Usage: "Your Radio Rig Infomation",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "output",
			Usage: "Your Output Setting Infomation",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "antenna",
			Usage: "Your Antenna Infomation",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "log_file_path",
			Usage: "Log File Path",
			Value: usr.HomeDir + "/hamradio_logger/log.json",
		}),
	}
	app := &cli.App{
		Commands: commands.CmdList,
		Flags:    flags,
		Before:   altsrc.InitInputSourceWithContext(flags, altsrc.NewYamlSourceFromFlagFunc("config")),
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
