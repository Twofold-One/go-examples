package basic

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

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

// example command to decode protol buffer file:
// $ sudo protoc --decode perftools.profiles.Profile /home/twofold_one/GitProjects/go/pprof/proto/profile.proto --proto_path /home/twofold_one/GitProjects/go/pprof/proto <  profile.pb
