package cli

import (
	"os"
	"os/exec"
)

func (s *service) CreatePackage() error {
	os.Chdir(s.folder)
	cmd := exec.Command("go", "mod", "init", s.packageName)
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
