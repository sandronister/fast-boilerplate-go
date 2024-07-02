package cli

type service struct {
	folder string
}

func NewService(folder string) *service {
	return &service{folder: folder}
}
