
prom:
	docker run \
        --network "host" \
        -p 9090:9090 \
        -v ./prometheus.yml:/etc/prometheus/prometheus.yml \
        prom/prometheus

run:
	go run main.go
