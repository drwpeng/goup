package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

const (
	DefaultHTTPErrorCode int = 600
)

// 别名
type (
	// interface
	Binder          = echo.Binder
	BindUnmarshaler = echo.BindUnmarshaler
	Context         = echo.Context
	Validator       = echo.Validator
	Renderer        = echo.Renderer
	Logger          = echo.Logger
	MiddlewareFunc  = echo.MiddlewareFunc
	HTTPError       = echo.HTTPError
	Map             = echo.Map

	// struct
	Group = echo.Group

)

var App *echo.Echo

var SelfPlugins []echo.MiddlewareFunc

func init() {
	App = echo.New()
	App.HTTPErrorHandler = HTTPErrorHandler
}

func HTTPErrorHandler(err error, c Context) {
	var (
		code = http.StatusInternalServerError
		msg  interface{}
	)

	if he, ok := err.(*HTTPError); ok {
		code = he.Code
		msg = he.Message
		if he.Internal != nil {
			err = fmt.Errorf("%v, %v", err, he.Internal)
		}
	} else if c.Echo().Debug {
		msg = err.Error()
	} else {
		msg = http.StatusText(code)
	}
	if v, ok := msg.(string); ok {
		msg = Map{"message": v}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead { // Issue #608
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, msg)
		}
		if err != nil {
			c.Logger().Error(err)
		}
	}
}

// NewHTTPError creates a new HTTPError instance.
func NewHTTPError(code int, message ...interface{}) *HTTPError {
	he := &HTTPError{Code: code, Message: http.StatusText(code)}
	if len(message) > 0 {
		he.Message = message[0]
	}
	return he
}

// 返回HTTPError，状态码统一、固定
func Error(message interface{}) *HTTPError {
	return NewHTTPError(DefaultHTTPErrorCode, message)
}

