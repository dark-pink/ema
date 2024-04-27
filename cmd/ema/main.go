package main

import (
	"darkp.ink/ema/tool"
	_ "darkp.ink/ema/tool/ema2json"
	_ "darkp.ink/ema/tool/json2ema"
)

func main() {
	tool.DoWithOS()
}
