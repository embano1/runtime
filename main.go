// Print runtime statistic and allow to burn CPU (defaults to "off")
package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

// Burns the CPU
func burn() {
	var i int64

	for {
		i++
	}
}

func tick() {
	tick := time.Tick(10 * time.Second)
	for range tick {
		fmt.Print(".")
	}
}

func main() {

	// Turn CPU burn off by default
	cpus := flag.Int("c", 0, "Numbers of Goroutines to spin up to burn CPU")
	flag.Parse()

	fmt.Println(time.Now().UTC())

	if *cpus > 0 {
		fmt.Printf("Starting %d Goroutines to burn CPU(s) without a deadline\n", *cpus)
		for i := 1; i <= *cpus; i++ {
			go burn()
		}
		go tick()
	} else {
		go func() {
			fmt.Println("CPU burning disabled, sleeping...")
			go tick()
		}()
	}

	// Print runtime statistics
	fmt.Println("Detected CPUs:", runtime.NumCPU())
	fmt.Println("Running Goroutines:", runtime.NumGoroutine())

	select {}
}
