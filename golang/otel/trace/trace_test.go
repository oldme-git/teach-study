package trace

import (
	"context"
	"log"
	"testing"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
)

func TestTraceHttp(t *testing.T) {
	ctx := context.Background()

	// 创建 OTLP HTTP 导出器，连接到 Jaeger
	exporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpointURL("http://srv.com:4318/v1/traces"))

	if err != nil {
		log.Fatalf("创建导出器失败: %v", err)
	}

	// 创建资源
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String("otel-traces-demo-http"),
		),
	)
	if err != nil {
		log.Fatalf("创建资源失败: %v", err)
	}

	// 创建 Tracer 提供器
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	// 设置全局 Tracer 提供器
	otel.SetTracerProvider(tp)

	// 创建一个新的 trace
	tracer := otel.Tracer("example-tracer")
	ctx, span := tracer.Start(ctx, "root-span")
	// 暂停 100ms
	time.Sleep(100 * time.Millisecond)
	// 结束 span
	span.End()

	// 创建子span
	_, childSpan := tracer.Start(ctx, "child-span")
	// 暂停 50ms
	time.Sleep(50 * time.Millisecond)
	childSpan.End()

	// 确保所有的 spans 都被发送
	if err := tp.Shutdown(ctx); err != nil {
		log.Fatalf("关闭 Tracer 提供器失败: %v", err)
	}
}

func TestTraceGrpc(t *testing.T) {
	ctx := context.Background()

	// 创建 OTLP gRPC 导出器，连接到 Jaeger
	exporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint("srv.com:4317"),
		otlptracegrpc.WithInsecure(),
	)

	if err != nil {
		log.Fatalf("创建导出器失败: %v", err)
	}

	// 创建资源
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String("otel-traces-demo-grpc"),
		),
	)
	if err != nil {
		log.Fatalf("创建资源失败: %v", err)
	}

	// 创建 Tracer 提供器
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	// 设置全局 Tracer 提供器
	otel.SetTracerProvider(tp)

	// 创建一个新的 trace
	tracer := otel.Tracer("example-tracer")
	ctx, span := tracer.Start(ctx, "root-span")
	// 暂停 100ms
	time.Sleep(100 * time.Millisecond)
	// 结束 span
	span.End()

	// 创建子span
	_, childSpan := tracer.Start(ctx, "child-span")
	// 暂停 50ms
	time.Sleep(50 * time.Millisecond)
	childSpan.End()

	// 确保所有的 spans 都被发送
	if err := tp.Shutdown(ctx); err != nil {
		log.Fatalf("关闭 Tracer 提供器失败: %v", err)
	}
}
