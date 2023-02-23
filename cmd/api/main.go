// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/tracer/stats"
	"github.com/cloudwego/hertz/pkg/common/utils"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"
	"github.com/hertz-contrib/registry/nacos"
	"mini_tiktok/cmd/api/rpc"
	"mini_tiktok/pkg/configs/config"
	"mini_tiktok/pkg/consts"
	nacosmw "mini_tiktok/pkg/nacos"
)

func Init() {
	// 配置初始化要放在最前面
	config.Init()
	nacosmw.Init()
	rpc.Init()
	hlog.SetLogger(hertzlogrus.NewLogger())
	hlog.SetLevel(hlog.LevelDebug)
}

func main() {
	Init()
	r := nacos.NewNacosRegistry(nacosmw.NacosClient)
	tracer, cfg := tracing.NewServerTracer()
	addr := "0.0.0.0:8080"
	h := server.New(
		server.WithHostPorts(addr),
		server.WithTraceLevel(stats.LevelDetailed),
		server.WithHandleMethodNotAllowed(true),
		server.WithRegistry(r, &registry.Info{
			ServiceName: consts.ApiServiceName,
			Addr:        utils.NewNetAddr("tcp", addr),
			Weight:      10,
			Tags:        nil,
		}),
		tracer,
		server.WithMaxRequestBodySize(100*1024*1024),
	)
	pprof.Register(h)
	h.Use(tracing.ServerMiddleware(cfg))
	register(h)
	h.Spin()
}
