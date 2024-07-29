package cli

import (
	"os"

	"github.com/sandronister/fast-bolierplate-go/internals/cli/files"
)

func (s *service) CreateDockerFile() error {

	file, err := os.Create(".Dockerfile")

	if err != nil {
		return err
	}

	defer file.Close()

	textFile := files.GetDockerFile()

	_, err = file.Write(textFile)

	if err != nil {
		return err
	}

	return nil
}
