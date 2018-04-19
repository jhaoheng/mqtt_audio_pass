package main

import (
  "log"
  "time"
)

func main() {

  for range time.Tick(60 * time.Second) {
    log.Println("..... debug mode hold")
  }
}
