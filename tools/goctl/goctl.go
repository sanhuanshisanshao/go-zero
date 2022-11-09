package main

import (
	"github.com/sanhuanshisanshao/go-zero/core/load"
	"github.com/sanhuanshisanshao/go-zero/core/logx"
	"github.com/sanhuanshisanshao/go-zero/tools/goctl/cmd"
)

func main() {
	logx.Disable()
	load.Disable()
	cmd.Execute()
}
