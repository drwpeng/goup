package v1

import (
	"myapp/pkg/rest"
	"myapp/internal/biz"
)

func init(){
	rest.Register(nil, "/myapp/v1/upload", &biz.Upload{})
}