package cli

import "os"

func (s *service) CreateInitialFolder() error {
	err := os.Mkdir(s.folder, 0755)
	if err != nil {
		return err
	}
	return nil
}
