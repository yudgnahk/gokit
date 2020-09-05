package utils

import (
	"strings"

	"github.com/khanghldk/gokit/constants"
)

type StandardizeParams struct {
	ModuleName     string
	BasePath       string
	ControllerName string
	ServiceName    string
	RepositoryName string
}

func StandardizedTemplate(content string, params StandardizeParams) string {
	content = strings.ReplaceAll(content, constants.BACK_STICK, constants.Backtick)
	content = strings.ReplaceAll(content, constants.MODULE_NAME, params.ModuleName)
	content = strings.ReplaceAll(content, constants.BASE_PATH, params.BasePath)

	content = strings.ReplaceAll(content, constants.CONTROLLER_NAME, Camel(params.ControllerName, false))
	content = strings.ReplaceAll(content, constants.SERVICE_NAME, Camel(params.ServiceName, false))
	content = strings.ReplaceAll(content, constants.REPOSITORY_NAME, Camel(params.RepositoryName, false))

	content = strings.ReplaceAll(content, constants.CONTROLLER_NAME_LOWER, Camel(params.ControllerName, true))
	content = strings.ReplaceAll(content, constants.SERVICE_NAME_LOWER, Camel(params.ServiceName, true))
	content = strings.ReplaceAll(content, constants.REPOSITORY_NAME_LOWER, Camel(params.RepositoryName, true))

	return content
}
