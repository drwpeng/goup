// +build wireinject

package main

import (
	"github.com/google/wire"
	"myapp/pkg/rest"
)

func InitMyApp() *rest.App {
	return wire.Build()
}