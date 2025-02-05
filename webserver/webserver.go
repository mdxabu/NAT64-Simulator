package webserver

import (
    "encoding/json"
    "fmt"
    "net/http"
    "sync"
)

var (
    mu                sync.Mutex  
    totalTranslations int
)

type Metrics struct {
    TotalTranslations int `json:"totalTranslations"`
}

func incrementTranslations() {
    mu.Lock()          
    totalTranslations++
    mu.Unlock()        
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()          // Lock the mutex before accessing the shared variable
    defer mu.Unlock()  // Ensure the mutex is unlocked after returning

    metrics := Metrics{
        TotalTranslations: totalTranslations,
    }

    w.Header().Set("Content-Type", "application/json")

    if err := json.NewEncoder(w).Encode(metrics); err != nil {
        http.Error(w, fmt.Sprintf("Error encoding metrics: %v", err), http.StatusInternalServerError)
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    html := `
    <html>
        <head>
            <title>NAT64 Gateway Simulator</title>
            <script>
                function fetchMetrics() {
                    fetch('/metrics')
                        .then(response => response.json())
                        .then(data => {
                            document.getElementById('totalTranslations').innerText = data.totalTranslations;
                        })
                        .catch(error => console.error('Error fetching metrics:', error));
                }

                // Fetch metrics every 1 second
                setInterval(fetchMetrics, 1000);  // Set the interval to 1 second
            </script>
        </head>
        <body>
            <h1>NAT64 Gateway Simulator</h1>
            <h2>Metrics:</h2>
            <p>Total Translations: <span id="totalTranslations">0</span></p>
        </body>
    </html>`

    // Serve the HTML page
    fmt.Fprintf(w, "%s", html)
}

func StartServer() {
    http.HandleFunc("/", homeHandler) 
    http.HandleFunc("/metrics", metricsHandler) 

    fmt.Println("Web server started at http://localhost:8080")
    go func() {
        if err := http.ListenAndServe(":8080", nil); err != nil {
            fmt.Println("Error starting web server:", err)
        }
    }()
}

func TranslateAndTrack(ipv6Addr string) {
    incrementTranslations()

    fmt.Printf("Translated IPv6 %s to IPv4\n", ipv6Addr)
}
