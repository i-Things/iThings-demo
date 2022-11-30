package main

import (
	"flag"
	"fmt"
	"github.com/i-Things/things/src/apisvr/apidirect"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/i4de/iThings-demo/src/papisvr/internal/config"
	"github.com/i4de/iThings-demo/src/papisvr/internal/handler"
	"github.com/i4de/iThings-demo/src/papisvr/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/papi.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.DisableStat()
	apiCtx := apidirect.NewApi(apidirect.ApiCtx{})
	ctx := svc.NewServiceContext(c, apiCtx)
	handler.RegisterHandlers(apiCtx.Server, ctx)

	apiCtx.Server.PrintRoutes()
	fmt.Printf("Starting papiSvr at %s:%d...\n", apiCtx.Svc.Config.Host, apiCtx.Svc.Config.Port)
	apiCtx.Server.Start()
	defer apiCtx.Server.Stop()
}
