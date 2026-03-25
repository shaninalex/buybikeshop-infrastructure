package main

import (
	"buybikeshop/apps/office/app/cmd"
	"os"
)

func main() {
	os.Exit(cmd.Execute())
}
