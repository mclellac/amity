package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mclellac/amity/lib/server"
	"github.com/mclellac/amity/lib/ui"

	"github.com/urfave/cli"
	"gopkg.in/gcfg.v1"
)

const (
	ServerConfigFile = "amityd.conf"
)

func readConfig(c *cli.Context) (server.Config, error) {
	configPath := c.GlobalString("config")
	config := server.Config{}

	if _, err := os.Stat(configPath); err != nil {
		return config, fmt.Errorf("%sERROR:%s Configuration file %s\"%s\"%s not found.", ui.Red, ui.Reset, ui.Red, configPath, ui.Reset)
	}

	err := gcfg.ReadFileInto(&config, ServerConfigFile)
	if err != nil {
		return config, err
	}

	return config, err
}

func main() {
	app := cli.NewApp()
	app.Name = "amityd"
	app.Usage = "(Amity Daemon) is the application server for the Amity service."
	app.Version = "0.0.5"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: ServerConfigFile,
			Usage: "specify an alternate server configuration file.",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start the server",
			Action: func(c *cli.Context) error {
				cfg, err := readConfig(c)
				if err != nil {
					log.Fatal(err)
				}

				d := server.Daemon{}

				if err = d.Run(cfg); err != nil {
					log.Fatal(err)
				}

				return nil
			},
		},
	}
	app.Run(os.Args)
}
