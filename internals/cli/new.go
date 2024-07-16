package cli

import (
	b64 "encoding/base64"
)

const configText = "cGFja2FnZSBjb25maWcKCmltcG9ydCAoCgkiZ2l0aHViLmNvbS9zcGYxMy92aXBlciIKKQoKdHlwZSBDb25mIHN0cnVjdCB7Cn0KCmZ1bmMgTG9hZENvbmZpZyhwYXRoIHN0cmluZykgKCpDb25mLCBlcnJvcikgewoJdmFyIGNmZyAqQ29uZgoJdmlwZXIuU2V0Q29uZmlnTmFtZSgiY29uZmlnIikKCXZpcGVyLlNldENvbmZpZ1R5cGUoImVudiIpCgl2aXBlci5BZGRDb25maWdQYXRoKHBhdGgpCgl2aXBlci5TZXRDb25maWdGaWxlKCIuZW52IikKCXZpcGVyLkF1dG9tYXRpY0VudigpCgoJZXJyIDo9IHZpcGVyLlJlYWRJbkNvbmZpZygpCglpZiBlcnIgIT0gbmlsIHsKCQlyZXR1cm4gbmlsLCBlcnIKCX0KCgllcnIgPSB2aXBlci5Vbm1hcnNoYWwoJmNmZykKCglpZiBlcnIgIT0gbmlsIHsKCQlyZXR1cm4gbmlsLCBlcnIKCX0KCglyZXR1cm4gY2ZnLCBuaWwKCn0K"

type service struct {
	folder      string
	packageName string
	config      []byte
}

func NewService(folder string, packageName string) *service {
	text, _ := b64.StdEncoding.DecodeString(configText)
	return &service{folder: folder, packageName: packageName, config: text}
}
