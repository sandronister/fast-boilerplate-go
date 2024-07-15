package cli

import (
	"os"
	"os/exec"
)

var library = "github.com/spf13/viper"

func (s *service) GetPackage() error {
	os.Chdir(s.folder)
	cmd := exec.Command("go", "get", library)
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
