package cli

import (
	"os"
)

func (s *service) CreatePrincipals() error {
	principals := []string{"internal", "pkg", "cmd", "config"}

	for _, principal := range principals {
		err := os.Mkdir(principal, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
