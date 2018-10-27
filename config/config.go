package config

type Config struct {
	Tasks []Task `yaml:"tasks"`
}

type Task struct {
	Name     string `yaml:"name"`
	Schedule string `yaml:"schedule"`
	Command  string `yaml:"command"`
}
