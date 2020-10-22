package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

func main() {

	tracer, closer := InitTracer("tracing-service")
	defer closer.Close()

	span := createMainSpan(tracer, "hello-span")
	defer span.Finish()

	doStub(span)

	logStub(span)

	// Play with context
	ctx := context.Background()
	ctxTracer, ctxCloser := InitTracer("context-service")
	defer ctxCloser.Close()
	opentracing.SetGlobalTracer(ctxTracer) //important!!!

	ctxSpan := createMainSpan(ctxTracer, "hello-context-span")
	defer ctxSpan.Finish()

	ctx = opentracing.ContextWithSpan(ctx, ctxSpan)

	doStubWithContext(ctx)

	logStubWithContext(ctx)
}

func createMainSpan(tracer opentracing.Tracer, name string) opentracing.Span {
	span := tracer.StartSpan(name)
	fmt.Println("Hello!!!")
	span.SetTag("http.url", "Localhost")
	span.SetTag("http.method", "GET")

	span.LogKV("event", "main")
	return span
}

func doStub(rootSpan opentracing.Span) {
	span := rootSpan.Tracer().StartSpan("doStub", opentracing.ChildOf(rootSpan.Context()))
	defer span.Finish()
	fmt.Println("Do Stub....")

	span.LogFields(log.String("event", "doStub"), log.Bool("success", true))

}

func logStub(rootSpan opentracing.Span) {
	span := rootSpan.Tracer().StartSpan("logStub", opentracing.ChildOf(rootSpan.Context()))
	defer span.Finish()
	fmt.Println("Log Stub....")

	span.LogFields(log.String("event", "logStub"), log.Bool("success", false))

}

func doStubWithContext(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "doStubWithContext")
	defer span.Finish()
	fmt.Println("Do Stub with context....")

	span.LogFields(log.String("event", "doStubWithContext"), log.Bool("success", true))

}

func logStubWithContext(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "logStubWithContext")

	defer span.Finish()
	fmt.Println("Log Stub with context....")

	span.LogFields(log.String("event", "logStubWithContext"), log.Bool("success", false))

}
