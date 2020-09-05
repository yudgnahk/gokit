package templates

// ErrorsTemplate ...
const ErrorsTemplate = `package errors

import (
	"fmt"
	"strconv"
)

// ErrorCode contains HTTP status, module and detail code.
type ErrorCode struct {
	status     int
	module     int
	detailCode int
}

func fmtErrorCode(status, module, code int) ErrorCode {
	return ErrorCode{
		status:     status,
		module:     module,
		detailCode: code,
	}
}

// Code returns the integer with format 4xxyyzz.
func (errCode ErrorCode) Code() int {
	errStr := fmt.Sprintf("%d%02d%02d", errCode.status, errCode.module, errCode.detailCode)
	code, _ := strconv.Atoi(errStr)
	return code
}

// Status returns HTTP status code.
func (errCode ErrorCode) Status() int {
	return errCode.status
}

// Module returns module error code.
func (errCode ErrorCode) Module() int {
	return errCode.module
}

// DetailCode returns detail error code.
func (errCode ErrorCode) DetailCode() int {
	return errCode.detailCode
}

// AppError describes application error.
type AppError struct {
	Meta          ErrorMeta 
	OriginalError error     
	ErrorCode     ErrorCode 
}

// ErrorMeta is the metadata of AppError.
type ErrorMeta struct {
	Code    int    
	Message string
}

func (appErr AppError) Error() string {
	if appErr.OriginalError != nil {
		return appErr.OriginalError.Error()
	}
	return appErr.Meta.Message
}

// New returns an AppError with args.
func New(errCode ErrorCode, args ...interface{}) error {
	return AppError{
		Meta: ErrorMeta{
			Code:    errCode.Code(),
			Message: GetErrorMessage(errCode, args...),
		},
		OriginalError: nil,
		ErrorCode:     errCode,
	}
}

// Newf returns an AppError with args and message.
func Newf(errCode ErrorCode, msg string, args ...interface{}) error {
	return AppError{
		Meta: ErrorMeta{
			Code:    errCode.Code(),
			Message: fmt.Sprintf(msg, args...),
		},
		OriginalError: nil,
		ErrorCode:     errCode,
	}
}

// Wrap returns an AppError with err, args.
func Wrap(errCode ErrorCode, err error, args ...interface{}) error {
	return AppError{
		Meta: ErrorMeta{
			Code:    errCode.Code(),
			Message: GetErrorMessage(errCode, args...),
		},
		OriginalError: err,
		ErrorCode:     errCode,
	}
}

// Wrapf returns an AppError with err, args and message.
func Wrapf(errCode ErrorCode, err error, msg string, args ...interface{}) error {
	return AppError{
		Meta: ErrorMeta{
			Code:    errCode.Code(),
			Message: fmt.Sprintf(msg, args...),
		},
		OriginalError: err,
		ErrorCode:     errCode,
	}
}
`

// ErrorCodesTemplate ...
const ErrorCodesTemplate = `package errors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Module constants definition.
const (
	ModuleCommon   = 00
)

// Common module error codes definition.
var (
	ErrValidation     = fmtErrorCode(http.StatusBadRequest, ModuleCommon, 1)
	ErrFieldInvalid   = fmtErrorCode(http.StatusBadRequest, ModuleCommon, 2)
	ErrUnauthorized   = fmtErrorCode(http.StatusUnauthorized, ModuleCommon, 1)
	ErrInternalServer = fmtErrorCode(http.StatusInternalServerError, ModuleCommon, 1)
	ErrNoResponse     = fmtErrorCode(http.StatusInternalServerError, ModuleCommon, 2)
	ErrPanicked       = fmtErrorCode(http.StatusInternalServerError, ModuleCommon, 3)
	ErrParseTime      = fmtErrorCode(http.StatusBadRequest, ModuleCommon, 4)
	ErrMaintain       = fmtErrorCode(http.StatusBadRequest, ModuleCommon, 5)
	ErrParse          = fmtErrorCode(http.StatusInternalServerError, ModuleCommon, 6)
)

type (
	translatedMessages struct {
		VI string 
		EN string
	}
	detailCodeMap map[string]translatedMessages
	statusMap     map[string]detailCodeMap
)

var errorMessageMap map[string]statusMap

func init() {
	errorMessageMap = make(map[string]statusMap)
}

// ErrorMessagesFilePath is path of error_messages.json file.
const ErrorMessagesFilePath = "./error_messages.json"

// InitErrorMessagesResource loads error messages resource.
func InitErrorMessagesResource() error {
	buf, err := ioutil.ReadFile(ErrorMessagesFilePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf, &errorMessageMap)
	if err != nil {
		return err
	}
	return nil
}

// GetErrorMessage gets error message from errorMessageMap.
func GetErrorMessage(errCode ErrorCode, args ...interface{}) string {
	msg := http.StatusText(errCode.Status())
	if errorMessageMap == nil {
		return msg
	}
	modules, ok := errorMessageMap[fmt.Sprintf("%02d", errCode.Module())]
	if modules == nil || !ok {
		return msg
	}
	statuses, ok := modules[fmt.Sprintf("%d", errCode.Status())]
	if statuses == nil || !ok {
		return msg
	}
	detailCodes, ok := statuses[fmt.Sprintf("%02d", errCode.DetailCode())]
	if !ok {
		return msg
	}
	if detailCodes.VI != "" {
		msg = detailCodes.EN
	}
	return fmt.Sprintf(msg, args...)
}
`