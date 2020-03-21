package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	// "net/http"
	// _ "net/http/pprof"
	"sync"
	"time"
)

//http://www.graphviz.org/Download_macos.php

// var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")

var memFile *os.File

func main() {
	flag.Parse()
	// if *cpuprofile != "" {
	//     f, err := os.Create(*cpuprofile)
	//     if err != nil {
	//         log.Fatal(err)
	//     }
	//     pprof.StartCPUProfile(f)
	//     defer pprof.StopCPUProfile()
	// }

	if *memprofile != "" {
		var err error
		memFile, err = os.Create(*memprofile)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("start write heap profile....")
			pprof.WriteHeapProfile(memFile)
			defer memFile.Close()
		}
	}

	// go func() {
	//     log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go work(&wg)
	}

	wg.Wait()
	// Wait to see the global run queue deplete.
	time.Sleep(300 * time.Second)
}

func work(wg *sync.WaitGroup) {
	time.Sleep(time.Second)

	var counter int
	for i := 0; i < 1e10; i++ {
		time.Sleep(time.Millisecond * 100)
		pprof.WriteHeapProfile(memFile)
		counter++
	}
	wg.Done()
}
