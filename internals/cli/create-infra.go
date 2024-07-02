package cli

import (
	"os"
)

func (s *service) CreateInfra() error {
	infra := []string{"internal/infra/database", "internal/infra/web", "internal/infra/handler"}

	for _, inf := range infra {
		err := os.MkdirAll(inf, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
