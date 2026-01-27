package scan

import (
    "fmt"
    "net"
    "sort"
    "sync"
    "time"
)

// ScanPorts scans a range of ports on a given hostname.
func ScanPorts(hostname string, startPort, endPort int) []int {
    var openPorts []int
    var wg sync.WaitGroup
    sem := make(chan struct{}, 1000) // Limit concurrency to 1000
    var mu sync.Mutex

    for port := startPort; port <= endPort; port++ {
        sem <- struct{}{}
        wg.Add(1)
        go func(p int) {
            defer wg.Done()
            address := fmt.Sprintf("%s:%d", hostname, p)
            conn, err := net.DialTimeout("tcp", address, 1500*time.Millisecond) // 1.5s timeout for speed
            if err == nil {
                conn.Close()
                mu.Lock()
                openPorts = append(openPorts, p)
                mu.Unlock()
            }
            <-sem
        }(port)
    }
    wg.Wait()
    sort.Ints(openPorts)
    return openPorts
}
