package main

import (
	"fmt"
	"github.com/uber/jaeger-client-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
	"io"
)

// initJaeger 将jaeger tracer设置为全局tracer
func initJaeger(service string) io.Closer {
	cfg := jaegerCfg.Configuration{
		// 将采样频率设置为1，每一个span都记录，方便查看测试结果
		Sampler: &jaegerCfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegerCfg.ReporterConfig{
			LogSpans: true,
			// 将span发往jaeger-collector的服务地址
			CollectorEndpoint: "http://192.168.10.43:14268/api/traces",
		},
	}
	closer, err := cfg.InitGlobalTracer(service, jaegerCfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return closer
}
