package main

import (
	_ "ShenMeYun-golang/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"ShenMeYun-golang/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
