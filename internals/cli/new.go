package cli

import "github.com/sandronister/fast-bolierplate-go/internals/cli/files"

type service struct {
	folder      string
	packageName string
	config      []byte
}

func NewService(folder string, packageName string) *service {
	text := files.GetConfigText()
	return &service{folder: folder, packageName: packageName, config: text}
}
