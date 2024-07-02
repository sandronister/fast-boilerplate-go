package cli

import (
	"os"
)

func (s *service) CreateInternal() error {
	internals := []string{"internal/entity", "internal/repository", "internal/infra", "internal/usecase"}

	for _, internal := range internals {
		err := os.MkdirAll(internal, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
