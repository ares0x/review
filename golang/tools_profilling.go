package main

import (
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create(*cpuprofile)
	if err != nil {

	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
}