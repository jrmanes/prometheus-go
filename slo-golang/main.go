package main

import (
    "net/http"
    "time"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var counter = promauto.NewCounter(prometheus.CounterOpts{
    Name: "orders_processed",
})

func main() {
    go func() {
        for {
            counter.Inc()
            // simulate some processing function
            time.Sleep(time.Second)
        }
    }()

    http.Handle("/metrics", promhttp.Handler())
    http.ListenAndServe(":2112", nil)
}

