package main

import (
	_ "goSimpleAdmin/internal/packed"

	_ "goSimpleAdmin/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"goSimpleAdmin/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
