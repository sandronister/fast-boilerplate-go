package cli

import (
	"os"

	"github.com/sandronister/fast-bolierplate-go/internals/cli/files"
)

func (s *service) CreateConfig() error {

	file, err := os.Create("config/config.go")

	if err != nil {
		return err
	}

	defer file.Close()

	configFile := files.GetConfigText()

	_, err = file.Write(configFile)

	if err != nil {
		return err
	}

	return nil
}
