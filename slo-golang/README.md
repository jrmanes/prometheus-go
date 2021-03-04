# Prometheus with Golang

In this repository you will find a Golang app which use prometheus to check the generated metrics.


# PromQL

```promql
promhttp_metric_handler_requests_total{code="200"}
promhttp_metric_handler_requests_total{code="500"}
promhttp_metric_handler_requests_total{code="503"} 
```


```promql
orders_processed{instance="127.0.0.1:2112"}
increase(orders_processed{instance="127.0.0.1:2112"}[5m])
rate(orders_processed{instance="127.0.0.1:2112"}[5m])
```

```promql
rate(orders_processed{instance="127.0.0.1:21112"}[5m])+
rate(orders_processed{instance="127.0.0.1:2112"}[5m])
```


## SLO-related queries

### API Errors Rate

Let's say we have counter http_requests_total that tracks the number of the received requests and has label status_code.


```promql
rate(http_requests_total{status_code=~"5.*"}[5m)
```

```promql
sum(rate(http_requests_total{status_code=~"5.*"}[5m]) 
/ sum(rate(http_requests_total[5m])
```

### Request Latency

Another metric response_latency_ms tracks the latency of the API responses in ms. We can check whether 95% of the responses take less or equal than 5 seconds.


```promql
histogram_quantile(
    0.95, 
    sum(rate(request_latency_ms_bucket[5m])) by (le)
) / 1e3 > 5
```

## Conclusion

As we can see Prometheus is a powerful and flexible tool. 
This repository shows Prometheus usage from basic instrumentation of a Go application to SLO-related PromQl queries.


---
Jose Ramón Mañes










