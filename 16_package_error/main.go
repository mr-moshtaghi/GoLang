package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type Config struct {
	// config fields go here
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, errors.Wrap(err, "can not open config file")
	}

	defer file.Close()

	config := Config{}

	// parse the file here

	return &config, nil
}

func main() {
	config, err := readConfig("/path/to/fake/file.toml")

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR : %s\n", err)
		os.Exit(1)
	}

	// normal operation

	fmt.Println(config)

}
