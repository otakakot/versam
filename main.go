package main

import (
	"runtime/debug"
)

var version string

func main() {
	if version == "" {
		if bi, ok := debug.ReadBuildInfo(); ok {
			version = bi.Main.Version
		}
	}

	println("versamp version ", version)
}
