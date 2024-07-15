package cli

import "os"

type service struct {
	folder      string
	packageName string
	config      []byte
}

func NewService(folder string, packageName string) *service {
	text, err := os.ReadFile("./config.txt")

	if err != nil {
		panic(err)
	}
	return &service{folder: folder, packageName: packageName, config: text}
}
