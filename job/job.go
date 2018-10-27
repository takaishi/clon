package job

import (
	"log"
	"os/exec"
	"strings"
)

type Job struct {
	Name    string
	Command string
}

func (c *Job) Run() {
	cmd := strings.Split(c.Command, " ")
	out, err := exec.Command(cmd[0], cmd[1:]...).Output()
	if err != nil {
		log.Printf("[ERROR] Failed to exec: %s\n", err)
	} else {
		log.Printf("[INFO] %s\n", out)
		log.Printf("[INFO] Success")
	}
}
