package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
    timeout := 3 * time.Second
    deadline := time.Now().Add(10 * time.Second)
    timeOutContext, _ := context.WithTimeout(context.Background(), timeout)
    cancelContext, cancel := context.WithCancel(context.Background())
    deadlineContext, _ := context.WithDeadline(context.Background(), deadline)


    go contextDemo(context.WithValue(timeOutContext, "name", "[timeoutContext]"))
    go contextDemo(context.WithValue(cancelContext, "name", "[cancelContext]"))
    go contextDemo(context.WithValue(deadlineContext, "name", "[deadlineContext]"))

    // Wait for the timeout to expire
    <- timeOutContext.Done()

    fmt.Println("Cancelling the cancel context...")
    cancel()

    <- cancelContext.Done()
    fmt.Println("The cancel context has been cancelled...")

    <- deadlineContext.Done()
    fmt.Println("The deadline context has been cancelled...")
}

func contextDemo(ctx context.Context) {
  deadline, ok := ctx.Deadline()
  name := ctx.Value("name")
  for {
    if ok {
      fmt.Println(name, "will expire at:", deadline)
    } else {
      fmt.Println(name, "has no deadline")
    }
    time.Sleep(time.Second)
  }
}
