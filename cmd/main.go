package main

import (
	"os"

	"github.com/chaewonkong/echo-demo/server"
)

func main() {
	code := server.Run()
	os.Exit(code)
}
