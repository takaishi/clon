package main

import (
	"github.com/robfig/cron"
	"github.com/takaishi/clon/config"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "/etc/clon.yml",
		},
	}

	app.Action = func(c *cli.Context) error {
		return action(c)
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func action(c *cli.Context) error {
	cfg, err := readConfig(c.String("config"))
	if err != nil {
		return err
	}

	l, err :=  time.LoadLocation("Local")
	if err != nil {
		return err
	}

	server := cron.NewWithLocation(l)

	for _, job := range cfg.Jobs {
		server.AddJob(job.Schedule, job)
	}

	server.Run()
	return nil
}

func readConfig(configPath string) (config.Config, error) {
	var cfg config.Config
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return cfg, err
	}
	err = yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
