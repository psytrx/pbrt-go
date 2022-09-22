package main

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

const (
	CPU_PROFILE = true
	MEM_PROFILE = true
)

func main() {
	if CPU_PROFILE {
		cpuProf, err := os.Create("./cpu.prof")
		if err != nil {
			log.Fatalf("could not create CPU profile: %s", err)
		}
		defer cpuProf.Close()

		if err := pprof.StartCPUProfile(cpuProf); err != nil {
			log.Fatalf("could not start CPU profile: %s", err)
		}
		defer pprof.StopCPUProfile()
	}

	start()

	if MEM_PROFILE {
		memProf, err := os.Create("./mem.prof")
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer memProf.Close()

		runtime.GC()
		if err := pprof.WriteHeapProfile(memProf); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
