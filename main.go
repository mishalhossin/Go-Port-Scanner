package main

import (
    "fmt"
    "net"
    "sync"
    "time"
)

func scanPort(host string, port int, timeout time.Duration, wg *sync.WaitGroup, results chan string) {
    fmt.Printf("Scanning port %v...\n", port)
    defer wg.Done()
    target := fmt.Sprintf("%s:%d", host, port)
    conn, err := net.DialTimeout("tcp", target, timeout)
    if err == nil {
        defer conn.Close()
        net.LookupPort("tcp", fmt.Sprintf("%d", port))
        results <- fmt.Sprintf("Port %v is open", port)
    }
}

func main() {
    var host string
    fmt.Print("Enter the target IP address: ")
    fmt.Scanln(&host)

    timeout := time.Second * 2

    fmt.Printf("Scanning ports on %s...\n", host)

    var wg sync.WaitGroup
    results := make(chan string)

    for port := 1; port <= 65535; port++ {
        wg.Add(1)
        go scanPort(host, port, timeout, &wg, results)
    }

    go func() {
        wg.Wait()
        close(results)
    }()

    for result := range results {
        fmt.Println(result)
    }
}
