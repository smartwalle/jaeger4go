package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/smartwalle/jaeger4go"
)

func main() {
	var cfg, err = jaeger4go.Load("./cfg.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	closer, err := cfg.InitGlobalTracer("test-service")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer closer.Close()

	span, ctx := opentracing.StartSpanFromContext(context.Background(), "1")
	span.LogKV("ss", "vv")
	span.Finish()

	span, _ = opentracing.StartSpanFromContext(ctx, "2")
	span.LogKV("vv", "ee")
	span.Finish()

}
