package utils

import (
	"strings"

	"github.com/khanghldk/gokit/constants"
)

// StandardizeParams ...
type StandardizeParams struct {
	ModuleName     string
	BasePath       string
	ControllerName string
	ServiceName    string
	RepositoryName string
}

// StandardizedTemplate ...
func StandardizedTemplate(content string, params StandardizeParams) string {
	content = strings.ReplaceAll(content, constants.BackStick, constants.Backtick)
	content = strings.ReplaceAll(content, constants.ModuleName, params.ModuleName)
	content = strings.ReplaceAll(content, constants.BasePath, params.BasePath)

	content = strings.ReplaceAll(content, constants.ControllerName, Camel(params.ControllerName, false))
	content = strings.ReplaceAll(content, constants.ServiceName, Camel(params.ServiceName, false))
	content = strings.ReplaceAll(content, constants.RepositoryName, Camel(params.RepositoryName, false))

	content = strings.ReplaceAll(content, constants.ControllerNameLower, Camel(params.ControllerName, true))
	content = strings.ReplaceAll(content, constants.ServiceNameLower, Camel(params.ServiceName, true))
	content = strings.ReplaceAll(content, constants.RepositoryNameLower, Camel(params.RepositoryName, true))

	return content
}
