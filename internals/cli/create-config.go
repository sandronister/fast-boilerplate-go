package cli

import "os"

func (s *service) CreateConfig() error {

	file, err := os.Create("config/config.go")

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(s.config)

	if err != nil {
		return err
	}

	return nil
}
