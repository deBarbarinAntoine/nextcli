package internal

import "embed"

const (
	componentTemplateFile = "templates/component.tmpl"
	componentTemplateName = "component"
)

//go:embed templates
var Templates embed.FS
