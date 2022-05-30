package basic

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

// TODO

// Difference between profiling and benchmark
// A benchmark collects runtime information about a specific function. Profiling is the collection of statistics for the whole program.
func ProfilingExample() {
	result := doSum()
	fmt.Println(result)

	// profiling itself
	// create file "profile.pb.gz"
	f, err := os.Create("/home/twofold_one/GitProjects/go/go-examples/basic/profile.pb.gz")
	if err != nil {
		log.Fatal(err)
	}
	// start CPU profiling and write profile result onto this file
	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatal(err)
	}
	// at the end of the main func we stop profiling
	defer pprof.StopCPUProfile()
}

func doSum() int {
	sum := 0
	for i := 0; i < 787766777; i++ {
		sum += 1
	}
	return sum
}
