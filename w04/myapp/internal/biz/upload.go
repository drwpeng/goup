package biz

import (
	"myapp/pkg/rest"
	"net/http"
)

type Upload struct {
	rest.Context
}

func (u *Upload) Post(c rest.Context) error {
	return c.JSON(http.StatusOK, "")
}