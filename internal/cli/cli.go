package cli

import "io"

type CLI struct {
	stdout io.Writer
}

type Config struct {
	Stdout io.Writer
}

func New(cfg *Config) *CLI {
	return &CLI{
		stdout: cfg.Stdout,
	}
}
