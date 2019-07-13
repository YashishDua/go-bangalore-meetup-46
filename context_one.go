package main

import (
	"context"
	"fmt"
	"time"
)

func main () {
  result := make(chan int)
  ctx, _ := context.WithCancel(context.Background())

  go Perform(ctx, result)

  <- time.After(time.Second * 3)
  //fmt.Println("Cancelation started")
  //cancel()

  code := <- result
  if code == 0 {
    fmt.Println("Call complete")
  } else {
    fmt.Println("Call canceled")
  }
}

func Perform(ctx context.Context, result chan int) {
  fmt.Println("Performing Task 1")

  select {
    case <- time.After(time.Second * 10):
      fmt.Println("Performing Task 2")
      result <- 0
    case <- ctx.Done():
      result <- 1
    }
}
