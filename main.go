package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

func main() {

	tracer, closer := InitTracer("tracing-service")
	defer closer.Close()

	span := createMainSpan(tracer)
	defer (*span).Finish()

	doStub(span)

	logStub(span)

}

func createMainSpan(tracer opentracing.Tracer) *opentracing.Span {
	span := tracer.StartSpan("hello-span")
	fmt.Println("Hello!!!")
	span.SetTag("http.url", "Localhost")
	span.SetTag("http.method", "GET")

	span.LogKV("event", "main")
	return &span
}

func doStub(rootSpan *opentracing.Span) {
	span := (*rootSpan).Tracer().StartSpan("doStub", opentracing.ChildOf((*rootSpan).Context()))
	defer span.Finish()
	fmt.Println("Do Stub....")

	span.LogFields(log.String("event", "doStub"), log.Bool("success", true))

}

func logStub(rootSpan *opentracing.Span) {
	span := (*rootSpan).Tracer().StartSpan("logStub", opentracing.ChildOf((*rootSpan).Context()))
	defer span.Finish()
	fmt.Println("Log Stub....")

	span.LogFields(log.String("event", "logStub"), log.Bool("success", false))

}
