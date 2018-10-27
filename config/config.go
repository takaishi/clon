package config

import "github.com/takaishi/clon/job"

type Config struct {
	Jobs []job.Job `yaml:"tasks"`
}
