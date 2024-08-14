package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	api "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
)

const meterName = "oldme_prometheus_testing"

var (
	requestHelloCounter      api.Int64Counter
	requestDurationHistogram api.Float64Histogram
)

func main() {
	ctx := context.Background()

	// 创建 prometheus 导出器
	exporter, err := prometheus.New()
	if err != nil {
		log.Fatal(err)
	}

	// 创建 meter
	provider := metric.NewMeterProvider(metric.WithReader(exporter))
	meter := provider.Meter(meterName)

	// 创建 counter 指标类型
	requestHelloCounter, err = meter.Int64Counter("requests_hello_total")
	if err != nil {
		log.Fatal(err)
	}

	// 创建 Histogram 指标类型
	requestDurationHistogram, err = meter.Float64Histogram(
		"request_hello_duration_seconds",
		api.WithDescription("记录 Hello 请求的耗时统计"),
		api.WithExplicitBucketBoundaries(0.1, 0.2, 0.5, 1, 2, 5),
	)
	if err != nil {
		log.Fatal(err)
	}

	go serveMetrics()
	go goroutineMock()

	ctx, _ = signal.NotifyContext(ctx, os.Interrupt)
	<-ctx.Done()
}

func serveMetrics() {
	log.Printf("serving metrics at localhost:2223/metrics")
	http.Handle("/metrics", promhttp.Handler())

	http.Handle("/index", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 记录 counter 指标
		requestHelloCounter.Add(r.Context(), 1)

		// 计算请求处理时间
		startTime := time.Now()
		// 模拟请求处理时间
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		defer func() {
			duration := time.Since(startTime).Seconds()
			requestDurationHistogram.Record(r.Context(), duration)
		}()

		_, _ = w.Write([]byte("Hello, Otel!"))
	}))

	err := http.ListenAndServe(":2223", nil) //nolint:gosec // Ignoring G114: Use of net/http serve function that has no support for setting timeouts.
	if err != nil {
		fmt.Printf("error serving http: %v", err)
		return
	}
}

// 随机模拟若干个协程
func goroutineMock() {
	for {
		go func() {
			// 等待若干秒
			var s = time.Duration(rand.Intn(10))
			time.Sleep(s * time.Second)
		}()
		time.Sleep(1 * time.Millisecond)
	}
}
