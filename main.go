package main

import (
    "fmt"
    "log"
    "net"
    "nat64-simulator/webserver" 
    "os"
    "bufio"
)

func main() {
    webserver.StartServer()

    file, err := os.Open("top-domains.txt")
    if err != nil {
        log.Fatalf("Error opening domains file: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        domain := scanner.Text()

        ipAddresses, err := net.LookupIP(domain)
        if err != nil {
            log.Printf("Error resolving domain %s: %v", domain, err)
            continue
        }

		fmt.Println("===============================")
        fmt.Printf("Domain: %s\n", domain)
        fmt.Println("Resolved IPs:")
        for _, ip := range ipAddresses {
            fmt.Println(ip.String())
            webserver.TranslateAndTrack(ip.String())  
			
        }
		fmt.Println("===============================")
        fmt.Println() 
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Error reading domains file: %v", err)
    }
}
