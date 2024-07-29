package cli

type service struct {
	folder      string
	packageName string
}

func NewService(folder string, packageName string) *service {
	return &service{folder: folder, packageName: packageName}
}
