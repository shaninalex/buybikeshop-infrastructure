package main

import (
	"buybikeshop/apps/datasource/app/cmd"
	"os"
)

func main() {
	os.Exit(cmd.Execute())
}
