package main

import (
	"buybikeshop/apps/warehouse/app/cmd"
	"os"
)

func main() {
	os.Exit(cmd.Execute())
}
