package main

import (
	"flag"
	"fmt"
	"time"
	"ultra-fast-port-scanner/pkg/report"
	"ultra-fast-port-scanner/pkg/scan"
)

func main() {
	start := time.Now()

	// Command line flags
	hostPtr := flag.String("host", "localhost", "Target hostname or IP")
	startPortPtr := flag.Int("start", 1, "Start port")
	endPortPtr := flag.Int("end", 1024, "End port")

	flag.Parse()

	fmt.Printf("Starting Ultra-Fast Port Scanner for %s (%d-%d)...\n", *hostPtr, *startPortPtr, *endPortPtr)

	// Perform scan
	openPorts := scan.ScanPorts(*hostPtr, *startPortPtr, *endPortPtr)

	// Print summary
	duration := time.Since(start)
	fmt.Printf("\nScan complete in %v. Found %d open ports.\n", duration, len(openPorts))

	// Print detailed report
	report.PrintReport(*hostPtr, openPorts)
}
