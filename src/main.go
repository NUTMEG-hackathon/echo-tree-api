package main

import (
  "github.com/mizu-ryo/echo-sam1/src/router"
)

func main() {
  router := router.NewRouter()
  router.Logger.Fatal(router.Start(":8888"))
}
