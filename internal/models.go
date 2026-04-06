//go:generate go run github.com/debarbarinantoine/go-enum-generate@latest --force
package internal

import "nextcli/internal/enum"

const (
	AppFolder          = "app"
	PageFuncName       = "NewPage"
	PageExtension      = ".tsx"
	ComponentExtension = ".tsx"
	LibExtension       = ".ts"
)

type NextFile struct {
	Name  string
	Type  enum.FileType
	Path  string
	IsLib bool
}
