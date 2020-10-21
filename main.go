package main

import (
	"fmt"
)

func main() {

	tracer, closer := InitTracer("tracing-service")
	defer closer.Close()

	span := tracer.StartSpan("hello-span")
	fmt.Println("Hello!!!")
	span.SetTag("http.url", "Localhost")
	span.SetTag("http.method", "GET")

	span.LogKV("event", "logging")

	span.Finish()
}
