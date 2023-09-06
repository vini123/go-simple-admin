package main

import (
	_ "goSimpleAdmin/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"goSimpleAdmin/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
