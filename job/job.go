package job

import (
	jp "github.com/kokardy/jpholiday"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Job struct {
	Name     string  `yaml:"name"`
	Schedule string  `yaml:"schedule"`
	Command  string  `yaml:"command"`
	Options  Options `yaml:"options"`
}

type Options struct {
	SkipJPHoliday bool `yaml:"skip_jp_holiday"`
}

func (j Job) Run() {
	now := time.Now()
	if j.Options.SkipJPHoliday && j.isJPHoliday(now) {
		log.Printf("[INFO] Skip to execute job: %s/%s/%s is Japanese holiday.", now.Year(), now.Month(), now.Day())
	} else {
		cmdStr := strings.Split(j.Command, " ")
		cmd := exec.Command(cmdStr[0], cmdStr[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			log.Printf("[ERROR] Failed to exec: %s\n", err)
		} else {
			log.Printf("[INFO] Success")
		}
	}
}

func (j Job) isJPHoliday(t time.Time) bool {
	day := jp.NewDate(t.Year(), t.Month(), t.Day())
	isJPHoliday, _ := day.Holiday()
	return isJPHoliday
}
