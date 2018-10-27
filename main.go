package main

import (
	"github.com/robfig/cron"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"

	"os/exec"
	"strings"
)

type Config struct {
	Tasks []Task `yaml:"tasks"`
}

type Task struct {
	Name     string `yaml:"name"`
	Schedule string `yaml:"schedule"`
	Command  string `yaml:"command"`
}

func (t *Task) Run() {
	cmd := strings.Split(t.Command, " ")
	out, err := exec.Command(cmd[0], cmd[1:]...).Output()
	if err != nil {
		log.Printf("[ERROR] Failed to exec: %s\n", err)
	} else {
		log.Printf("[INFO] %s\n", out)
		log.Printf("[INFO] Success")
	}
}

type clonJob struct {
	name string
	cmd  string
}

func (c *clonJob) Run() {
	cmd := strings.Split(c.cmd, " ")
	out, err := exec.Command(cmd[0], cmd[1:]...).Output()
	if err != nil {
		log.Printf("[ERROR] Failed to exec: %s\n", err)
	} else {
		log.Printf("[INFO] %s\n", out)
		log.Printf("[INFO] Success")
	}
}

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "config, c",
		},
	}

	app.Action = func(c *cli.Context) error {
		var cfg Config
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
			j := &clonJob{name: task.Name, cmd: task.Command}
			server.AddJob(task.Schedule, j)
		}
		for _, entry := range server.Entries() {
			log.Printf("[DEBUG] entry: %+v\n", entry)
			log.Printf("[DEBUG] entry.Job: %+v\n", entry.Job)
		}
		server.Run()

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
