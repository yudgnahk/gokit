package templates

// BaseController ...
const BaseController = `package controllers

import (
"github.com/gin-gonic/gin"
"MODULE_NAME/utils"
)

type Base struct{}

// JSON responds a HTTP request with JSON data.
func (h *Base) JSON(c *gin.Context, data interface{}) {
	utils.JSON(c, data)
}

// HandleError handles error of HTTP request.
func (h *Base) HandleError(c *gin.Context, err error) {
	utils.HandleError(c, err)
}`

// ControllerTemplate ...
const ControllerTemplate = `package controllers

import (
	"github.com/gin-gonic/gin"
	"MODULE_NAME/services"
)

// CONTROLLER_NAMEController handles all request of CONTROLLER_NAME controller.
type CONTROLLER_NAMEController struct {
	Base
	controller_nameService services.CONTROLLER_NAMEService
}

// NewCONTROLLER_NAMEController returns a new instance of CONTROLLER_NAMEController.
func NewCONTROLLER_NAMEController(controller_nameService services.CONTROLLER_NAMEService) *CONTROLLER_NAMEController {
	return &CONTROLLER_NAMEController{
		controller_nameService: controller_nameService,
	}
}

// TODO: Implement DoSomething
func (c *CONTROLLER_NAMEController) DoSomething(ctx *gin.Context) {
	c.JSON(ctx, "something")
}
`

// AppController ...
const AppController = `package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"MODULE_NAME/dtos"
)

// AppController handles all request of app module.
type AppController struct {
	Base
}

// NewAppController returns a new instance of AppController.
func NewAppController() *AppController {
	return &AppController{}
}

// Health handles health check.
// @Summary Health
// @Description Handles health check API
// @Tags App
// @Accept json
// @Produce json
// @Success 200 {object} dtos.HealthResponse
// @Failure 500 {object} errors.AppError
// @Router /health [get]
func (c *AppController) Health(ctx *gin.Context) {
	c.JSON(ctx, &dtos.HealthResponse{Meta: dtos.Meta{
		Code:    http.StatusOK,
		Message: "OK",
	}})
}
`