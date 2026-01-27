# Ultra-Fast Port Scanner

A high-performance port scanner written in Go using massive concurrency (Goroutines).

## Features

- **Massive Concurrency**: Scans thousands of ports in seconds using lightweight Goroutines.
- **Worker Pool**: Efficiently manages resources with a semaphore-based worker pool (default 1000 workers).
- **Clean Output**: Displays results in a formatted table.

## Installation

1.  Clone the repository:
    ```bash
    git clone <repository-url>
    cd ultra-fast-port-scanner
    ```
2.  Build the project (requires Go 1.21+):
    ```bash
    go build -o scanner cmd/scanner/main.go
    ```

## Usage

Run the compiled binary:

```bash
./scanner -host example.com -start 1 -end 1024
```

### Options

- `-host`: Target hostname or IP (default: `localhost`).
- `-start`: Start port (default: `1`).
- `-end`: End port (default: `1024`).

## Performance

Capable of scanning 1000 ports in under 3 seconds on standard network conditions.

## License

Copyright (c) 27 Janvier 2026 - Antigravity
See [LICENSE](LICENSE) for details.
