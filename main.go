package main

import (
	"github.com/robfig/cron"
	"github.com/takaishi/clon/config"
	"github.com/takaishi/clon/job"
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
			Name: "config, c",
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
	var cfg config.Config
	log.Printf("[DEBUG] action")
	data, err := ioutil.ReadFile(c.String("config"))
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] %+v\n", cfg)

	server := cron.New()

	for _, task := range cfg.Tasks {
		log.Printf("[DEBUG] %+v\n", task)
		j := &job.Job{Name: task.Name, Command: task.Command}
		server.AddJob(task.Schedule, j)
	}
	for _, entry := range server.Entries() {
		log.Printf("[DEBUG] entry: %+v\n", entry)
		log.Printf("[DEBUG] entry.Job: %+v\n", entry.Job)
	}
	server.Run()
	return nil
}
