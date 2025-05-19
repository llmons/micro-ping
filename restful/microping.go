package main

import (
	"flag"
	"fmt"

	"micro-ping/restful/internal/base/response"
	"micro-ping/restful/internal/config"
	"micro-ping/restful/internal/handler"
	"micro-ping/restful/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/microping-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetOkHandler(response.OkHandler)
	httpx.SetErrorHandlerCtx(response.ErrHandler(c.Name))

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
