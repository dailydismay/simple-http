package configs

import (
	"github.com/jessevdk/go-flags"
	"os"
)

type Config struct {
	LogLevel string `long:"log_level" env:"LOG_LEVEL" description:"log level for logrus" required:"yes"`
	Port string`long:"port" env:"PORT" description:"port to bind" required:"yes"`
}

func Parse() (*Config, error) {
	var config Config

	p := flags.NewParser(&config, flags.HelpFlag|flags.PassDoubleDash)

	if _, err := p.ParseArgs(os.Args); err != nil {
		return nil, err
	}

	return &config, nil
}
