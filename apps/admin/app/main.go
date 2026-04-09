package main

import (
	"buybikeshop/apps/admin/app/cmd"
	"os"
)

func main() {
	os.Exit(cmd.Execute())
}
