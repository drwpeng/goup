package rest

import (
	"github.com/labstack/echo"
)

// 所有HTTP方法的接口，主要用于简化类型转换
type (
	GET interface {
		Get(c echo.Context) error
	}
	POST interface {
		Post(c echo.Context) error
	}
	PUT interface {
		Put(c echo.Context) error
	}
	DELETE interface {
		Delete(c echo.Context) error
	}

	HEAD interface {
		Head(c echo.Context) error
	}

	PATCH interface {
		Patch(c echo.Context) error
	}

	OPTIONS interface {
		Options(c echo.Context) error
	}

	// Resourcer interface {
	// 	GetName() string
	// 	SetName(name string) string
	// }
)

// 资源
type Resource struct {
	Name string
}

func (self *Resource) GetName() string {
	return self.Name
}

func (self *Resource) SetName(name string) error {
	self.Name = name
	return nil
}

// 注册路由
func Register(g *echo.Group, uri string, c interface{}) {
	// c.(Resourcer).SetName(uri)
	if g == nil {
		g = App.Group("")
	}
	if m, ok := c.(GET); ok {
		g.GET(uri, m.Get)
	}
	if m, ok := c.(POST); ok {
		g.POST(uri, m.Post)
	}
	if m, ok := c.(PUT); ok {
		g.PUT(uri, m.Put)
	}
	if m, ok := c.(DELETE); ok {
		g.DELETE(uri, m.Delete)
	}
	if m, ok := c.(HEAD); ok {
		g.HEAD(uri, m.Head)
	}
	if m, ok := c.(PATCH); ok {
		g.PATCH(uri, m.Patch)
	}
	if m, ok := c.(OPTIONS); ok {
		g.OPTIONS(uri, m.Options)
	}
}
