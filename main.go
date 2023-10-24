package main

import (
	"orb/cmd"
	_ "orb/cmd/controller"
)

func main() {
	cmd.Execute()
}
