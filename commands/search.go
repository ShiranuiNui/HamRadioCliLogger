package commands

import (
	"log"
	"os"
	"os/user"

	repo "github.com/ShiranuiNui/HamRadioCliLogger/repositories"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func init() {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:     "callsign",
			Usage:    "callsign of worked station",
			Aliases:  []string{"c"},
			Required: true,
		},
	}
	CmdList = append(CmdList, &cli.Command{
		Name:  "search",
		Usage: "Search QSO Logging by Callsign",
		Flags: flags,
		Before: altsrc.InitInputSource(flags, func() (altsrc.InputSourceContext, error) {
			usr, _ := user.Current()
			if _, err := os.Stat(usr.HomeDir + "/.config/hamradio_logger"); os.IsNotExist(err) {
				os.Mkdir(usr.HomeDir+"/.config/hamradio_logger", 0777)
			}
			if _, err := os.Stat(usr.HomeDir + "/.config/hamradio_logger/config.yaml"); os.IsNotExist(err) {
				return &altsrc.MapInputSource{}, nil
			}
			return altsrc.NewYamlSourceFromFile(usr.HomeDir + "/.config/hamradio_logger/config.yaml")
		}),
		Action: func(c *cli.Context) error {
			qso, err := repo.ReadQSOByCallsign(c.String("callsign"))
			if err != nil {
				log.Fatal(err)
				return cli.Exit("", 1)
			}
			repo.ShowQSOTable(qso)
			return nil
		}})
}
