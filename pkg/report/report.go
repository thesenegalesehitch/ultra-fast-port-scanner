package report

import (
    "fmt"
    "os"
    "text/tabwriter"
)

// PrintReport formats and prints the scan results to the console.
func PrintReport(hostname string, openPorts []int) {
    w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
    fmt.Fprintf(w, "PORT\tSTATUS\n")
    for _, p := range openPorts {
        fmt.Fprintf(w, "%d/tcp\tOPEN\n", p)
    }
    w.Flush()
}
