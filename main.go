package main

import (

	_ "test/internal/packed"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gogf/gf/v2/os/gctx"
	"test/internal/cmd"
	"test/internal/service"
)

func main() {

	service.Init()
	cmd.Main.Run(gctx.New())
}
