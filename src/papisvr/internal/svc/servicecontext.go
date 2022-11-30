package svc

import (
	"github.com/i-Things/things/src/apisvr/apidirect"
	"github.com/i4de/iThings-demo/src/papisvr/internal/config"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config     config.Config
	CheckToken rest.Middleware
}

func NewServiceContext(c config.Config, ctx apidirect.ApiCtx) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		CheckToken: ctx.Svc.CheckToken,
	}
}
