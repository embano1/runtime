// Print runtime statistic and allow to burn CPU (defaults to "off")
package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

func burn() {
	var i int64

	for {
		i++
	}
}

func main() {

	// Turn CPU burn off by default
	cpus := flag.Int("c", 0, "Numbers of Goroutines to spin up to burn CPU")
	flag.Parse()

	if *cpus > 0 {
		fmt.Printf("Starting %d Goroutines to burn CPU(s) without a deadline\n", *cpus)
		for i := 1; i <= *cpus; i++ {
			go burn()
		}
	} else {
		go func() {
			fmt.Println("CPU burning disabled, sleeping for 1 hour before exiting")
			time.Sleep(1 * time.Hour)
		}()
	}

	// Print runtime statistics
	fmt.Println("Detected CPUs:", runtime.NumCPU())
	fmt.Println("Running Goroutines:", runtime.NumGoroutine())

	select {}
}
